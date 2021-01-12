package view

import (
	"reflect"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type View struct {
	Application    *gtk.Application
	MainWindow     *gtk.ApplicationWindow `gtk:"main_window"`
	SettingsWindow *gtk.Window            `gtk:"settings_window"`

	MainHeaderBar     *gtk.HeaderBar `gtk:"main_header_bar"`
	SettingsHeaderBar *gtk.HeaderBar `gtk:"settings_header_bar"`

	ButtonStack *gtk.Stack `gtk:"button_stack"`
	StatusStack *gtk.Stack `gtk:"status_stack"`

	StatusOkImage   *gtk.Image `gtk:"status_ok_image"`
	StatusFailImage *gtk.Image `gtk:"status_fail_image"`

	StatusSpinner *gtk.Spinner `gtk:"status_spinner"`

	VerifyButton   *gtk.Button `gtk:"verify_button"`
	CancelButton   *gtk.Button `gtk:"cancel_button"`
	SaveButton     *gtk.Button `gtk:"save_button"`
	SettingsButton *gtk.Button `gtk:"settings_button"`

	FileChooserButton *gtk.FileChooserButton `gtk:"file_chooser_button"`
	HashValueEntry    *gtk.Entry             `gtk:"hash_value_entry"`

	ErrorDialog *gtk.MessageDialog `gtk:"error_dialog"`
}

func New(appID string) *View {
	// Construct view
	view := new(View)

	// Define new signal emitted when widgets are initialized
	glib.SignalNew("ready")

	view.Application, _ = gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	view.Application.Connect("activate", func() {
		// Load UI from resources
		builder, err := gtk.BuilderNewFromResource("/data/checksumo.ui")
		if err != nil {
			panic(err)
		}

		// Get widgets from UI definition
		viewStruct := reflect.ValueOf(view).Elem()
		for i := 0; i < viewStruct.NumField(); i++ {
			field := viewStruct.Field(i)
			structField := viewStruct.Type().Field(i)
			widget := structField.Tag.Get("gtk")

			if widget == "" {
				continue
			}

			obj, err := builder.GetObject(widget)
			if err != nil {
				panic(err)
			}

			field.Set(reflect.ValueOf(obj).Convert(field.Type()))
		}

		// Widgets are ready now
		view.Application.Emit("ready")

		// Show and add main window
		view.MainWindow.Present()
		view.Application.AddWindow(view.MainWindow)
	})

	return view
}
