package views

import (
	"path/filepath"

	"github.com/dawidd6/checksumo/src/constants"
	"github.com/dawidd6/checksumo/src/settings"
	"github.com/dawidd6/checksumo/src/utils"

	"github.com/dawidd6/checksumo/src/presenters"

	"github.com/gotk3/gotk3/glib"

	"github.com/gotk3/gotk3/gtk"
)

type View interface {
	Activate(*gtk.Application)
}

type view struct {
	MainWindow        *gtk.ApplicationWindow `gtk:"main_window"`
	MainHeaderBar     *gtk.HeaderBar         `gtk:"main_header_bar"`
	ButtonStack       *gtk.Stack             `gtk:"button_stack"`
	VerifyButton      *gtk.Button            `gtk:"verify_button"`
	CancelButton      *gtk.Button            `gtk:"cancel_button"`
	FileChooserButton *gtk.FileChooserButton `gtk:"file_chooser_button"`
	HashValueEntry    *gtk.Entry             `gtk:"hash_value_entry"`
	ErrorDialog       *gtk.MessageDialog     `gtk:"error_dialog"`
	ResultOkDialog    *gtk.MessageDialog     `gtk:"result_ok_dialog"`
	ResultFailDialog  *gtk.MessageDialog     `gtk:"result_fail_dialog"`
}

func New() View {
	return &view{}
}

func (view *view) Activate(app *gtk.Application) {
	// Bind widgets
	utils.BindWidgets(view, constants.UIResourcePath)

	// Create presenter
	presenter := presenters.New(view)

	view.FileChooserButton.Connect("file-set", presenter.SetFile)
	view.HashValueEntry.Connect("changed", presenter.SetHash)
	view.HashValueEntry.Connect("activate", presenter.StartHashing)
	view.VerifyButton.Connect("clicked", presenter.StartHashing)
	view.CancelButton.Connect("clicked", presenter.StopHashing)

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

func (view *view) notify(title, body string) {
	if !settings.ShowNotifications() {
		return
	}

	notification := glib.NotificationNew(title)
	notification.SetBody(body)
	notification.SetDefaultAction("app.bring-up")

	app, _ := view.MainWindow.GetApplication()
	app.SendNotification(app.GetApplicationID(), notification)
}

func (view *view) GetFile() string {
	return view.FileChooserButton.GetFilename()
}

func (view *view) GetHash() string {
	hash, _ := view.HashValueEntry.GetText()
	return hash
}

func (view *view) OnResultError(err error) {
	if !view.MainWindow.IsActive() {
		title, _ := view.ErrorDialog.GetProperty("text")
		view.notify(title.(string), err.Error())
	}
	view.ErrorDialog.FormatSecondaryText(err.Error())
	view.ErrorDialog.Run()
	view.ErrorDialog.Hide()
}

func (view *view) OnResultSuccess() {
	if !view.MainWindow.IsActive() {
		text, _ := view.ResultOkDialog.GetProperty("text")
		filePath := view.FileChooserButton.GetFilename()
		fileName := filepath.Base(filePath)
		view.notify(text.(string), fileName)
	}
	view.ResultOkDialog.Run()
	view.ResultOkDialog.Hide()
}

func (view *view) OnResultFailure() {
	if !view.MainWindow.IsActive() {
		text, _ := view.ResultFailDialog.GetProperty("text")
		filePath := view.FileChooserButton.GetFilename()
		fileName := filepath.Base(filePath)
		view.notify(text.(string), fileName)
	}
	view.ResultFailDialog.Run()
	view.ResultFailDialog.Hide()
}

func (view *view) OnFileOrHashSet(isReady bool, hashType string) {
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

func (view *view) OnProgressUpdate(progress float64) {
	view.HashValueEntry.SetProgressFraction(progress)
}

func (view *view) OnProcessStart() {
	view.ButtonStack.SetVisibleChild(view.CancelButton)
	view.FileChooserButton.SetSensitive(false)
	view.HashValueEntry.SetSensitive(false)
}

func (view *view) OnProcessStop() {
	view.ButtonStack.SetVisibleChild(view.VerifyButton)
	view.FileChooserButton.SetSensitive(true)
	view.HashValueEntry.SetSensitive(true)
}
