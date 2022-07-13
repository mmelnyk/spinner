package spinner

import (
	"io"
	"time"
)

type optionNoOutput struct {
}

func (*optionNoOutput) apply(spinner *spinner) {
	spinner.output = io.Discard
}

type optionDuration struct {
	delay time.Duration
}

func (opt *optionDuration) apply(spinner *spinner) {
	spinner.delay = opt.delay
}

type optionStyle struct {
	style string
}

func (opt *optionStyle) apply(spinner *spinner) {
	spinner.runes = []rune(opt.style)
}

type optionShowElapsedTimer struct {
}

func (opt *optionShowElapsedTimer) apply(spinner *spinner) {
	spinner.showelapsed = true
}
