package view

import (
	"reflect"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type View struct {
	Application *gtk.Application

	MainWindow     *gtk.ApplicationWindow `gtk:"main_window"`
	SettingsWindow *gtk.Window            `gtk:"settings_window"`

	MainHeaderBar     *gtk.HeaderBar `gtk:"main_header_bar"`
	SettingsHeaderBar *gtk.HeaderBar `gtk:"settings_header_bar"`

	ButtonStack *gtk.Stack `gtk:"button_stack"`

	VerifyButton   *gtk.Button `gtk:"verify_button"`
	CancelButton   *gtk.Button `gtk:"cancel_button"`
	SaveButton     *gtk.Button `gtk:"save_button"`
	SettingsButton *gtk.Button `gtk:"settings_button"`

	FileChooserButton *gtk.FileChooserButton `gtk:"file_chooser_button"`
	HashValueEntry    *gtk.Entry             `gtk:"hash_value_entry"`

	ErrorDialog      *gtk.MessageDialog `gtk:"error_dialog"`
	ResultOkDialog   *gtk.MessageDialog `gtk:"result_ok_dialog"`
	ResultFailDialog *gtk.MessageDialog `gtk:"result_fail_dialog"`
}

func New(appName, appID, localeDomain, localeDir, uiResource string) *View {
	view := new(View)

	// Initialize localization
	glib.InitI18n(localeDomain, localeDir)

	view.Application, _ = gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	view.Application.Connect("activate", func() {
		// Load UI from resources
		builder, err := gtk.BuilderNewFromResource(uiResource)
		if err != nil {
			panic(err)
		}

		// Get widgets from UI definition
		vStruct := reflect.ValueOf(view).Elem()
		for i := 0; i < vStruct.NumField(); i++ {
			field := vStruct.Field(i)
			structField := vStruct.Type().Field(i)
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

		// Don't delete settings window after closing it, just hide
		view.SettingsWindow.HideOnDelete()

		// Show and add main window
		view.MainWindow.Present()
		view.Application.AddWindow(view.MainWindow)
	})

	return view
}
