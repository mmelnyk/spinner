package spinner

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

func NewSpinner(ctx context.Context, opts ...Option) *Spinner {
	spinner := &Spinner{
		ctx:    ctx,
		delay:  time.Duration(150 * time.Millisecond),
		runes:  []rune(StyleDot),
		output: os.Stdout,
	}

	for _, opt := range opts {
		opt.apply(spinner)
	}

	return spinner
}

const (
	StylePipe       = "|/-\\"
	StyleDot        = "⠁⠂⠄⡀⢀⠠⠐⠈"
	Style3Dots      = "⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏"
	StyleHDots      = "⢹⢺⢼⣸⣇⡧⡗⡏"
	StyleCompexDots = "⠁⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈⠈"
	StyleBars       = "-=≡="
	StyleHBar       = "▉▊▋▌▍▎▏▎▍▌▋▊▉"
	StyleVBar       = "▁▃▄▅▆▇█▇▆▅▄▃"
	StyleRounds     = "◜◝◞◟"
	StyleGlobe      = "🌍🌎🌏"
	StyleMoon       = "🌑🌒🌓🌔🌕🌖🌗🌘"
	StyleClock      = "🕐🕑🕒🕓🕔🕕🕖🕗🕘🕙🕚🕛"
)

type Spinner struct {
	once     sync.Once       // run only once
	ctx      context.Context // ctx is a context
	delay    time.Duration   // delay is how often the indicator should be updated
	nextrune int             // next rune to show
	runes    []rune          // runes holds the chosen sequece of symbols
	output   io.Writer       // output holds output destination
	message  string          // process description
}

type Option interface {
	apply(*Spinner)
}

type optionDuration struct {
	delay time.Duration
}

func (opt *optionDuration) apply(spinner *Spinner) {
	spinner.delay = opt.delay
}

func WithDelay(delay time.Duration) Option {
	return &optionDuration{delay: delay}
}

type optionStyle struct {
	style string
}

func (opt *optionStyle) apply(spinner *Spinner) {
	spinner.runes = []rune(opt.style)
}

func WithStyle(style string) Option {
	return &optionStyle{style: style}
}

func (spin *Spinner) Process(message string) {
	spin.once.Do(func() {
		spin.message = message
		go func() {
		loop:
			for {
				select {
				case <-spin.ctx.Done():
					break loop
				case <-time.After(spin.delay):
					spin.update()
				}
			}
			spin.done()
		}()
	})
}

func (spin *Spinner) update() {
	fmt.Fprintf(spin.output, "\r%c %s", spin.runes[spin.nextrune], spin.message)
	spin.nextrune += 1
	if spin.nextrune >= len(spin.runes) {
		spin.nextrune = 0
	}
}

func (spin *Spinner) done() {
	// do nothing in this version
}
