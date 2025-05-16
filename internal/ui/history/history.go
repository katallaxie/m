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

	go history.onUpdate()

	return history
}

func (h *History) onUpdate() {
	pctx := h.Application.Context()

	for {
		select {
		case <-pctx.Context().Done():
			return
		case change := <-h.Application.GetStore().Subscribe():
			treeRoot := tview.NewTreeNode("📚 Chats")
			curr := change.Curr()

			for _, chat := range curr.History.Chats {
				node := tview.NewTreeNode(chat.Name).
					SetReference(chat.ID).
					SetColor(tcell.ColorLightCoral).
					SetSelectable(true)
				treeRoot.AddChild(node)
				h.SetCurrentNode(node)
			}

			h.SetRoot(treeRoot)
		}
	}
}

func (h *History) onInputCapture(event *tcell.EventKey) *tcell.EventKey {
	return event
}
