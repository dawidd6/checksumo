package main

// #cgo pkg-config: gio-2.0
// #include "resources.h"
import "C"

import (
	"os"
	"reflect"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var (
	AppName   string
	AppID     string
	LocaleDir string
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

func NewView() *View {
	view := new(View)
	presenter := NewPresenter(view)

	// Initialize localization
	glib.InitI18n(AppName, LocaleDir)

	view.Application, _ = gtk.ApplicationNew(AppID, glib.APPLICATION_FLAGS_NONE)
	view.Application.Connect("activate", func() {
		// Load UI from resources
		builder, err := gtk.BuilderNewFromResource("/data/checksumo.ui")
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

		// Connect handlers to signals when widgets are ready
		view.FileChooserButton.Connect("file-set", presenter.SetFile)
		view.HashValueEntry.Connect("changed", presenter.SetHash)
		view.HashValueEntry.Connect("activate", presenter.StartHashing)
		view.VerifyButton.Connect("clicked", presenter.StartHashing)
		view.CancelButton.Connect("clicked", presenter.StopHashing)
		view.SettingsButton.Connect("clicked", view.SettingsWindow.Present)
		view.SaveButton.Connect("clicked", view.SettingsWindow.Hide)

		// Show and add main window
		view.MainWindow.Present()
		view.Application.AddWindow(view.MainWindow)
	})

	return view
}

func (view *View) Show() int {
	return view.Application.Run(os.Args)
}
