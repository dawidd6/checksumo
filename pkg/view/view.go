package view

import (
	"log"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type View struct {
	Application       *gtk.Application
	ApplicationWindow *gtk.ApplicationWindow
	VerifyButton      *gtk.Button
	HashEntry         *gtk.Entry
	FileButton        *gtk.FileChooserButton
	VerifyingSpinner  *gtk.Spinner
	ErrorDialog       *gtk.MessageDialog
	StatusStack       *gtk.Stack
	StatusOkImage     *gtk.Image
	StatusFailImage   *gtk.Image
	HashLabel         *gtk.Label
	signals           map[string]interface{}
}

func New(appID string) *View {
	var (
		view View
		err  error
	)

	view.Application, err = gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatal(err)
	}

	_, err = view.Application.Application.Connect("activate", func() {
		builder, err := gtk.BuilderNewFromFile("ui.glade")
		if err != nil {
			log.Fatal(err)
		}

		builder.ConnectSignals(view.signals)

		obj, err := builder.GetObject("main_window")
		if err != nil {
			log.Fatal(err)
		}
		view.ApplicationWindow = obj.(*gtk.ApplicationWindow)

		obj, err = builder.GetObject("hash_entry")
		if err != nil {
			log.Fatal(err)
		}
		view.HashEntry = obj.(*gtk.Entry)

		obj, err = builder.GetObject("file_button")
		if err != nil {
			log.Fatal(err)
		}
		view.FileButton = obj.(*gtk.FileChooserButton)

		obj, err = builder.GetObject("verify_button")
		if err != nil {
			log.Fatal(err)
		}
		view.VerifyButton = obj.(*gtk.Button)

		obj, err = builder.GetObject("verifying_spinner")
		if err != nil {
			log.Fatal(err)
		}
		view.VerifyingSpinner = obj.(*gtk.Spinner)

		obj, err = builder.GetObject("error_dialog")
		if err != nil {
			log.Fatal(err)
		}
		view.ErrorDialog = obj.(*gtk.MessageDialog)

		obj, err = builder.GetObject("status_stack")
		if err != nil {
			log.Fatal(err)
		}
		view.StatusStack = obj.(*gtk.Stack)

		obj, err = builder.GetObject("status_ok_image")
		if err != nil {
			log.Fatal(err)
		}
		view.StatusOkImage = obj.(*gtk.Image)

		obj, err = builder.GetObject("status_fail_image")
		if err != nil {
			log.Fatal(err)
		}
		view.StatusFailImage = obj.(*gtk.Image)

		obj, err = builder.GetObject("hash_label")
		if err != nil {
			log.Fatal(err)
		}
		view.HashLabel = obj.(*gtk.Label)

		view.ApplicationWindow.ShowAll()
		view.Application.AddWindow(view.ApplicationWindow)
	})
	if err != nil {
		log.Fatal(err)
	}

	return &view
}

func (view *View) SetSignals(signals map[string]interface{}) {
	view.signals = signals
}
