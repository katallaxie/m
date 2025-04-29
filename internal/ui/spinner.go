package ui

import (
	"sync"
	"time"

	"github.com/derailed/tview"
	"github.com/gdamore/tcell/v2"
)

type Spinner struct {
	tview.TextView // or anything
	loopOnce       sync.Once
	stop           chan struct{}
}

func (s *Spinner) Draw(screen tcell.Screen) {
	s.loopOnce.Do(func() {
		s.stop = make(chan struct{})

		go func() {
			ticker := time.NewTicker(time.Second) // any duration
			defer ticker.Stop()

			for {
				select {
				case <-s.stop:
					return
				case <-ticker.C:
					// Call self with screen.
					s.Draw(screen)
				}
			}
		}()
	})

	// Draw the spinner
}
