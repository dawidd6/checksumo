package views

import (
	"path/filepath"

	"github.com/dawidd6/checksumo/src/settings"

	"github.com/dawidd6/checksumo/src/constants"
	"github.com/dawidd6/checksumo/src/utils"

	"github.com/dawidd6/checksumo/src/presenters"

	"github.com/gotk3/gotk3/glib"

	"github.com/gotk3/gotk3/gtk"
)

type MainView interface {
	Activate(*gtk.Application)
}

type mainView struct {
	MainWindow        *gtk.ApplicationWindow
	MainHeaderBar     *gtk.HeaderBar
	SettingsButton    *gtk.Button
	ButtonStack       *gtk.Stack
	VerifyButton      *gtk.Button
	CancelButton      *gtk.Button
	FileChooserButton *gtk.FileChooserButton
	HashValueEntry    *gtk.Entry
	ErrorDialog       *gtk.MessageDialog
	ResultOkDialog    *gtk.MessageDialog
	ResultFailDialog  *gtk.MessageDialog
}

func NewMainView() MainView {
	return &mainView{}
}

func (view *mainView) Activate(app *gtk.Application) {
	// Bind widgets
	utils.BindWidgets(view, constants.UIResourcePath)

	// Create presenter
	presenter := presenters.New(view)

	// Connect handlers to events
	view.FileChooserButton.Connect("file-set", presenter.SetFile)
	view.HashValueEntry.Connect("changed", presenter.SetHash)
	view.HashValueEntry.Connect("activate", presenter.StartHashing)
	view.VerifyButton.Connect("clicked", presenter.StartHashing)
	view.CancelButton.Connect("clicked", presenter.StopHashing)
	view.SettingsButton.Connect("clicked", NewSettingsView().Activate)

	// Show main window
	view.MainWindow.Present()

	// Create actions
	bringUpAction := glib.SimpleActionNew("bring-up", nil)
	bringUpAction.Connect("activate", view.MainWindow.Present)
	quitAction := glib.SimpleActionNew("quit", nil)
	quitAction.Connect("activate", app.Quit)

	// Set keyboard shortcuts for actions
	app.SetAccelsForAction("app.quit", []string{"<Ctrl>Q"})

	// Add actions
	app.AddAction(bringUpAction)
	app.AddAction(quitAction)

	// Add main window
	app.AddWindow(view.MainWindow)
}

func (view *mainView) notify(title, body string) {
	notification := glib.NotificationNew(title)
	notification.SetBody(body)
	notification.SetDefaultAction("app.bring-up")

	app, _ := view.MainWindow.GetApplication()
	app.SendNotification(app.GetApplicationID(), notification)
}

func (view *mainView) GetFile() string {
	return view.FileChooserButton.GetFilename()
}

func (view *mainView) GetHash() string {
	hash, _ := view.HashValueEntry.GetText()
	return hash
}

func (view *mainView) OnResultError(err error) {
	if view.MainWindow.IsActive() || !settings.ShowNotifications() {
		view.ErrorDialog.FormatSecondaryText(err.Error())
		view.ErrorDialog.Run()
		view.ErrorDialog.Hide()
	} else {
		title, _ := view.ErrorDialog.GetProperty("text")
		view.notify(title.(string), err.Error())
	}
}

func (view *mainView) OnResultSuccess() {
	if view.MainWindow.IsActive() || !settings.ShowNotifications() {
		view.ResultOkDialog.Run()
		view.ResultOkDialog.Hide()
	} else {
		text, _ := view.ResultOkDialog.GetProperty("text")
		filePath := view.FileChooserButton.GetFilename()
		fileName := filepath.Base(filePath)
		view.notify(text.(string), fileName)
	}
}

func (view *mainView) OnResultFailure() {
	if view.MainWindow.IsActive() || !settings.ShowNotifications() {
		view.ResultFailDialog.Run()
		view.ResultFailDialog.Hide()
	} else {
		text, _ := view.ResultFailDialog.GetProperty("text")
		filePath := view.FileChooserButton.GetFilename()
		fileName := filepath.Base(filePath)
		view.notify(text.(string), fileName)
	}
}

func (view *mainView) OnFileOrHashSet(isReady bool, hashType string) {
	if view.MainHeaderBar.GetSubtitle() != hashType {
		view.MainHeaderBar.SetSubtitle(hashType)
	}
	if view.VerifyButton.GetSensitive() != isReady {
		view.VerifyButton.SetSensitive(isReady)
	}
	if view.HashValueEntry.GetProgressFraction() > 0 {
		view.HashValueEntry.SetProgressFraction(0.0)
	}
}

func (view *mainView) OnProgressUpdate(progress float64) {
	view.HashValueEntry.SetProgressFraction(progress)
}

func (view *mainView) OnProcessStart() {
	view.ButtonStack.SetVisibleChild(view.CancelButton)
	view.FileChooserButton.SetSensitive(false)
	view.HashValueEntry.SetSensitive(false)
}

func (view *mainView) OnProcessStop() {
	view.ButtonStack.SetVisibleChild(view.VerifyButton)
	view.FileChooserButton.SetSensitive(true)
	view.HashValueEntry.SetSensitive(true)
}
