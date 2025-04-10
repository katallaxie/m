package chat

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/katallaxie/m/internal/store"
	"github.com/katallaxie/m/internal/ui"
	"github.com/rivo/tview"
)

// NotebookList is a list of notebooks.
type NotebookList[S store.State] struct {
	app ui.Application[S]
	*tview.TreeView
}

// NewNotebookList returns a new notebook list.
func NewNotebookList[S store.State](app ui.Application[store.State]) *NotebookList[S] {
	notebookList := &NotebookList[S]{
		TreeView: tview.NewTreeView(),
	}

	treeRoot := tview.NewTreeNode("📚 Library")
	notebookList.SetRoot(treeRoot)
	notebookList.SetCurrentNode(treeRoot)
	notebookList.SetTitle(" 📚 Notebooks ")
	notebookList.SetBorder(true)

	notebookList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			fmt.Printf("tab pressed\n")
		}

		return event
	})

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

				notebookList.SetRoot(treeRoot)
			})
		}
	}()

	return notebookList
}
