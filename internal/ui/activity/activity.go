package activity

import (
	"time"

	"github.com/katallaxie/m/internal/state"
	"github.com/katallaxie/m/internal/ui"
	"github.com/navidys/tvxwidgets"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Activity is a chat activity.
type Activity struct {
	*tview.Flex
	gauge *tvxwidgets.ActivityModeGauge
	stop  chan struct{}
	app   ui.Application[state.State]
}

// NewActivity returns a new chat activity.
func NewActivity(app ui.Application[state.State]) *Activity {
	activity := &Activity{
		Flex:  tview.NewFlex(),
		gauge: tvxwidgets.NewActivityModeGauge(),
		app:   app,
		stop:  make(chan struct{}),
	}

	activity.gauge.SetPgBgColor(tcell.ColorOrange)

	activity.SetTitle(" 🏃‍♂️ Activity")
	activity.SetRect(10, 4, 50, 3)
	activity.SetBorder(true)

	activity.SetDirection(tview.FlexColumn)
	sub := app.GetStore().Subscribe()

	go func() {
		for current := range sub {
			if current.Status == state.Loading {
				activity.Clear()
				activity.AddItem(activity.gauge, 0, 1, true)
				go activity.onLoading()
			}

			if current.Status != state.Loading {
				activity.stop <- struct{}{}
				activity.Clear()
			}
		}
	}()

	return activity
}

func (activity *Activity) onLoading() {
	tick := time.NewTicker(500 * time.Millisecond)
	defer tick.Stop()

	for {
		select {
		case <-activity.stop:
			activity.Clear()
			return
		case <-tick.C:
			activity.gauge.Pulse()
			activity.app.Draw()
		}
	}
}
