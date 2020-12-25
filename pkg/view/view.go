package view

import (
	"log"
        "reflect"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type View struct {
	Application       *gtk.Application
	ApplicationWindow *gtk.ApplicationWindow `gtk:"main_window"` // TODO
	VerifyButton      *gtk.Button
	HashEntry         *gtk.Entry
	FileButton        *gtk.FileChooserButton
	VerifyingSpinner  *gtk.Spinner
	ErrorDialog       *gtk.MessageDialog
	AboutDialog       *gtk.AboutDialog
	StatusStack       *gtk.Stack
	StatusOkImage     *gtk.Image
	StatusFailImage   *gtk.Image
	HashLabel         *gtk.Label
	AboutButton       *gtk.ModelButton

	signals map[string]interface{}
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
		builder, err := gtk.BuilderNewFromResource("/data/ui.glade")
		if err != nil {
			log.Fatal(err)
		}

		builder.ConnectSignals(view.signals)

                viewStruct := reflect.ValueOf(&view).Elem()
	        for i := 0; i < viewStruct.NumField(); i++ {
		    field := viewStruct.Field(i)
		    widget := viewStruct.Type().Field(i).Tag.Get("gtk")
		    if widget == "" {
			continue
		    }

		    obj, err := builder.GetObject(widget)
		    if err != nil {
			log.Fatal(err)
		    }

		    field.Set(reflect.ValueOf(obj).Convert(field.Type()))
	        }

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
