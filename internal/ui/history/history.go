package history

import (
	"github.com/gdamore/tcell/v2"
	"github.com/katallaxie/m/internal/store"
	"github.com/katallaxie/m/internal/ui"
	"github.com/rivo/tview"
)

// History is a list of chats
type History struct {
	Application ui.Application[store.State]
	*tview.TreeView
}

// NewHistory creates a new history view
func NewHistory(app ui.Application[store.State]) *History {
	history := &History{
		Application: app,
		TreeView:    tview.NewTreeView(),
	}

	treeRoot := tview.NewTreeNode("📚 Chats")
	history.SetRoot(treeRoot)
	history.SetCurrentNode(treeRoot)
	history.SetTitle(" 📚 History ")
	history.SetBorder(true)
	history.SetInputCapture(history.onInputCapture)

	sub := app.GetStore().Subscribe()

	go func() {
		for change := range sub {
			app.QueueUpdateDraw(func() {
				treeRoot := tview.NewTreeNode("📚 Library")

				for _, notebook := range change.Curr().Notebooks {
					node := tview.NewTreeNode(notebook.Name).
						SetReference(notebook.ID).
						SetColor(tcell.ColorLightCoral).
						SetSelectable(true)
					treeRoot.AddChild(node)
				}

				history.SetRoot(treeRoot)
			})
		}
	}()

	return history
}

func (h *History) onInputCapture(event *tcell.EventKey) *tcell.EventKey {
	return event
}
