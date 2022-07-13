package spinner

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"go.melnyk.org/mansi"
)

type spinner struct {
	delay       time.Duration // delay is how often the indicator should be updated
	nextrune    int           // next rune to show
	runes       []rune        // runes holds the chosen sequece of symbols
	output      io.Writer     // output holds output destination
	started     time.Time     // when the spinner started
	showelapsed bool          // show elapsed timer
	message     string        // process description
	changed     bool          // indicate if process description is changed
	mu          sync.Mutex    // rw mutex for a message
}

func (spin *spinner) Process(ctx context.Context) {
	if spin.output == io.Discard {
		return
	}

	go func() {
		spin.started = time.Now()
		spin.update()
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			case <-time.After(spin.delay):
				spin.update()
			}
		}
		spin.done()
	}()
}

func (spin *spinner) Message(message string) {
	spin.mu.Lock()
	defer spin.mu.Unlock()

	if spin.message != message { // update only new process description
		spin.message = message
		spin.changed = true
	}
}

func (spin *spinner) update() {
	var elapsed string

	spin.mu.Lock()
	defer spin.mu.Unlock()
	if spin.showelapsed {
		elapsed = " (" + time.Since(spin.started).Round(time.Second).String() + ")"
	}

	fmt.Fprintf(spin.output, "\r%c %s%s%s", spin.runes[spin.nextrune], spin.message, elapsed, mansi.LineEraseToEnd)
	spin.changed = false
	spin.nextrune += 1
	if spin.nextrune >= len(spin.runes) {
		spin.nextrune = 0
	}
}

func (spin *spinner) done() {
	spin.mu.Lock()
	changed := spin.changed
	spin.mu.Unlock()

	if changed {
		spin.update()
	}
}
