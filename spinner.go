package spinner

import (
	"context"
	"os"
	"time"
)

type Spinner interface {
	Process(ctx context.Context)
	Message(msg string)
}

type Option interface {
	apply(*spinner)
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

func WithoutOutput() Option {
	return &optionNoOutput{}
}

func WithDelay(delay time.Duration) Option {
	return &optionDuration{delay: delay}
}

func WithStyle(style string) Option {
	return &optionStyle{style: style}
}

func WithElapsedTimer() Option {
	return &optionShowElapsedTimer{}
}

func NewSpinner(opts ...Option) Spinner {
	spinner := &spinner{
		delay:  time.Duration(150 * time.Millisecond),
		runes:  []rune(StyleDot),
		output: os.Stdout,
	}

	for _, opt := range opts {
		opt.apply(spinner)
	}

	return spinner
}
