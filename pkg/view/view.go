package view

import (
	"reflect"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type View struct {
	Application       *gtk.Application
	ApplicationWindow *gtk.ApplicationWindow `gtk:"main_window"`

	HeaderBar *gtk.HeaderBar `gtk:"header_bar"`

	StatusStack     *gtk.Stack   `gtk:"status_stack"`
	OkImage         *gtk.Image   `gtk:"ok_image"`
	FailImage       *gtk.Image   `gtk:"fail_image"`
	ProgressSpinner *gtk.Spinner `gtk:"progress_spinner"`

	ButtonStack  *gtk.Stack  `gtk:"button_stack"`
	VerifyButton *gtk.Button `gtk:"verify_button"`
	CancelButton *gtk.Button `gtk:"cancel_button"`

	FileChooserButton *gtk.FileChooserButton `gtk:"file_chooser_button"`
	HashValueEntry    *gtk.Entry             `gtk:"hash_value_entry"`

	ErrorDialog *gtk.MessageDialog `gtk:"error_dialog"`
}

func New(appID, version string) *View {
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
			widget := viewStruct.Type().Field(i).Tag.Get("gtk")
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
		view.ApplicationWindow.Present()
		view.Application.AddWindow(view.ApplicationWindow)
	})

	return view
}
