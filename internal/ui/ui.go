package ui

import (
	"os"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/katallaxie/m/internal/entity"
	"github.com/katallaxie/m/internal/ui/chat"
	"github.com/katallaxie/m/internal/ui/help"
	"github.com/katallaxie/m/internal/ui/infobar"
	"github.com/katallaxie/m/internal/ui/utils"

	"github.com/epiclabs-io/winman"
	"github.com/rivo/tview"
)

type ComponentLayout struct {
	MenuList     *tview.List
	BookmarkList *tview.TreeView
	LogList      *tview.TextArea
	OutputPanel  InitOutputPanelComponents
}

type UI struct {
	App    *tview.Application
	WinMan *winman.Manager
	Layout *ComponentLayout

	currentPage string
	pages       *tview.Pages
	menu        *tview.TextView
	chat        *chat.Chat
	help        *help.Help
	infoBar     *infobar.InfoBar

	Theme *entity.Theme
}

func (u *UI) SetFocus(p tview.Primitive) {
	go u.App.QueueUpdateDraw(func() {
		u.App.SetFocus(p)
	})
}

func (u UI) Run() error {
	u.App.EnableMouse(true)
	return u.App.Run()
}

func (u UI) QuitApplication() {
	u.App.Stop()
}

func NewUI(session entity.Session) UI {
	app := tview.NewApplication()
	wm := winman.NewWindowManager()

	ui := UI{
		App:     app,
		WinMan:  wm,
		Theme:   &entity.TerminalTheme,
		pages:   tview.NewPages(),
		infoBar: infobar.NewInfoBar("M", "0.1.0"),
		chat:    chat.NewChat("M", "0.1.0"),
	}

	ui.help = help.NewHelp("M", "0.1.0")

	// menu items
	menuItems := [][]string{
		{utils.HelpScreenKey.Label(), ui.help.GetTitle()},
		{utils.ChatScreenKey.Label(), ui.chat.GetTitle()},
		// {utils.SystemScreenKey.Label(), app.system.GetTitle()},
		// {utils.PodsScreenKey.Label(), app.pods.GetTitle()},
		// {utils.ContainersScreenKey.Label(), app.containers.GetTitle()},
		// {utils.VolumesScreenKey.Label(), app.volumes.GetTitle()},
		// {utils.ImagesScreenKey.Label(), app.images.GetTitle()},
		// {utils.NetworksScreenKey.Label(), app.networks.GetTitle()},
		// {utils.SecretsScreenKey.Label(), app.secrets.GetTitle()},
	}
	ui.menu = newMenu(menuItems)

	ui.pages.AddPage(ui.help.GetTitle(), ui.help, true, false)
	ui.pages.AddPage(ui.chat.GetTitle(), ui.chat, true, false)

	window := wm.NewWindow().
		Show().
		SetRoot(ui.setupAppLayout()).
		SetBorder(false)

	window.Maximize()
	app.SetRoot(wm, true)

	// listen for user input
	ui.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == utils.AppExitKey.Key {
			app.Stop()
			os.Exit(0)
		}

		event = utils.ParseKeyEventKey(event)

		// previous and next screen keys
		switch event.Rune() {
		case utils.NextScreenKey.Rune():
			return nil
		default:
		}

		// normal page key switch
		switch event.Key() { //nolint:exhaustive
		case utils.HelpScreenKey.EventKey():
			ui.switchToScreen(ui.help.GetTitle())

			return nil
		case utils.ChatScreenKey.EventKey():
			ui.switchToScreen(ui.chat.GetTitle())

			return nil
		}

		return event
	})

	ui.startupSequence()
	return ui
}

func (u *UI) startupSequence() {
	// u.loadStartupUI()
	// u.loadBookmarks()
	// u.startLogDumper()
	// u.startArgsConnection()
}

// setupAppLayout sets up the main grid layout of the application.
func (u *UI) setupAppLayout() *tview.Flex {

	// Setup the main layout
	// splitSidebar := tview.NewFlex().SetDirection(tview.FlexRow).
	// 	AddItem(u.Layout.MenuList, 15, 1, true).
	// 	AddItem(u.Layout.BookmarkList, 0, 1, false)

	// splitMainPanel := tview.NewFlex().SetDirection(tview.FlexRow).
	// 	AddItem(u.pages, 0, 3, false).
	// 	AddItem(u.Layout.LogList, 0, 1, false)

	// childLayout := tview.NewFlex().SetDirection(tview.FlexColumn).
	// 	AddItem(splitSidebar, 35, 1, true).
	// 	AddItem(splitMainPanel, 0, 4, false)

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(u.infoBar, 3, 1, false).
		AddItem(u.pages, 0, 1, false).
		AddItem(u.menu, 1, 1, false)

	return layout
}

type InitOutputPanelComponents struct {
	Layout   *tview.Flex
	TextArea *tview.TextArea
	Buffer   string
}

// InitOutputPanel initializes the output panel on the main screen
func (u *UI) InitOutputPanel() InitOutputPanelComponents {
	output := tview.NewTextArea()
	output.SetWrap(false)
	output.SetMaxLength(1)
	// output.SetTextStyle(tcell.StyleDefault.
	// 	Foreground(tcell.ColorGreen))

	// u.initOutputPanel_handleTextArea(output)

	layout := tview.NewFlex()
	layout.SetDirection(tview.FlexRow)
	layout.SetBorder(true)
	layout.SetTitle(" Output ")
	layout.AddItem(output, 0, 1, true)
	// layout.AddItem(u.initOutputPanel_PanelBar(), 1, 1, false)

	return InitOutputPanelComponents{
		Layout:   layout,
		TextArea: output,
		Buffer:   "",
	}
}

// InitBookmarkMenu initializes the bookmark sidebar menu
func (u *UI) InitBookmarkMenu() *tview.TreeView {
	treeRoot := tview.NewTreeNode("📚 Library")
	bookmarkList := tview.NewTreeView().
		SetRoot(treeRoot).
		SetCurrentNode(treeRoot)

	bookmarkList.SetBorder(true)
	bookmarkList.SetBorderPadding(1, 1, 1, 1)
	bookmarkList.SetTitle(" 📚 Notebooks ")

	// u.InitBookmarkMenu_SetInputCapture(bookmarkList)
	// u.InitBookmarkMenu_SetSelection(bookmarkList)

	return bookmarkList
}

// InitLogList initializes the log panel on the main screen
func (u *UI) InitLogList() *tview.TextArea {
	logPanel := tview.NewTextArea()
	// logPanel.SetDynamicColors(true)
	logPanel.SetTitle(" 📃 Prompt ")
	logPanel.SetBorder(true)
	logPanel.SetWordWrap(true)
	logPanel.SetBorderPadding(1, 1, 1, 1)
	logPanel.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			u.PrintOutput(entity.Output{
				WithHeader: true,
				Content:    "Hello World",
			})

			logPanel.SetText("", false)
		}

		return event
	})

	// u.InitLogList_SetInputCapture(logPanel)

	return logPanel
}

func (u *UI) InitLogList_SetInputCapture(logPanel *tview.TextView) {
	// inputField := tview.NewInputField().
	// 	SetLabel("Enter a number: ").
	// 	SetPlaceholder("E.g. 1234").
	// 	SetFieldWidth(10).
	// 	SetAcceptanceFunc(tview.InputFieldInteger).
	// 	SetDoneFunc(func(key tcell.Key) {
	// 		app.Stop()
	// 	})

	// logPanel.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	// 	switch event.Key() {
	// 	case tcell.KeyTAB:
	// 		u.App.SetFocus(u.Layout.MenuList)
	// 	}
	// 	return event
	// })
}

// PrintOutput used to print output to the output panel
func (u *UI) PrintOutput(param entity.Output) {
	var (
		// metadata  string
		newBuffer string
	)

	out := u.Layout.OutputPanel
	_, _, width, _ := out.TextArea.GetRect()

	timeHeader := time.Now().Format("15:04:05 02/01/2006")

	if param.WithHeader {
		payloadHeader := strings.Repeat(string(tcell.RuneCkBoard), 2) + "[ 🧑‍💻 User ]" + (strings.Repeat(string(tcell.RuneCkBoard), width-46)) + "[ " + timeHeader + " ]" + strings.Repeat(string(tcell.RuneCkBoard), 2) + "\n\n"
		newBuffer += payloadHeader

		responseHeader := "\n\n" + strings.Repeat(string(tcell.RuneCkBoard), 2) + "[ 🤖 Assistant ]" + (strings.Repeat(string(tcell.RuneCkBoard), width-47)) + "[ " + timeHeader + " ]" + strings.Repeat(string(tcell.RuneCkBoard), 2) + "\n"
		newBuffer += responseHeader + param.Content
	} else {
		newBuffer = param.Content
	}

	out.TextArea.SetText(newBuffer, param.CursorAtEnd)
}
