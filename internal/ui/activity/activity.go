package activity

import (
	"time"

	"github.com/katallaxie/m/internal/store"
	"github.com/katallaxie/m/internal/ui"
	"github.com/katallaxie/pkg/utilx"
	"github.com/navidys/tvxwidgets"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Activity is a chat activity.
type Activity struct {
	*tview.Flex
	gauge *tvxwidgets.ActivityModeGauge
	stop  chan struct{}
	app   ui.Application[store.State]
}

// NewActivity returns a new chat activity.
func NewActivity(app ui.Application[store.State]) *Activity {
	activity := &Activity{
		Flex:  tview.NewFlex(),
		gauge: tvxwidgets.NewActivityModeGauge(),
		app:   app,
		stop:  make(chan struct{}),
	}

	activity.gauge.SetPgBgColor(tcell.ColorOrange)

	activity.SetTitle(" 🏃‍♂️ Activity")
	activity.SetRect(10, 4, 50, 4)
	activity.SetBorder(true)

	activity.SetDirection(tview.FlexColumn)
	sub := app.GetStore().Subscribe()

	go func() {
		for change := range sub {
			if utilx.Equal(change.Curr().Status, change.Prev().Status) {
				continue
			}

			if change.Curr().Status == store.Loading {
				activity.Clear()
				activity.AddItem(activity.gauge, 0, 1, true)
				go activity.onLoading()
			}

			if change.Curr().Status != store.Loading {
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
