package chat

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// NotebookList is a list of notebooks.
type NotebookList struct {
	*tview.TreeView
}

// NewNotebookList returns a new notebook list.
func NewNotebookList() *NotebookList {
	notebookList := &NotebookList{
		TreeView: tview.NewTreeView(),
	}

	treeRoot := tview.NewTreeNode("📚 Library")
	notebookList.SetRoot(treeRoot)
	notebookList.SetCurrentNode(treeRoot)
	notebookList.SetTitle(" 📚 Notebooks Library ")
	notebookList.SetBorder(true)

	notebookList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			fmt.Printf("tab pressed\n")
		}

		return event
	})

	return notebookList
}
