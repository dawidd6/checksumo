package view

import (
	"log"
	"reflect"

	"github.com/gotk3/gotk3/gdk"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type View struct {
	Application       *gtk.Application
	ApplicationWindow *gtk.ApplicationWindow `gtk:"main_window"`
	VerifyButton      *gtk.Button            `gtk:"verify_button"`
	HashEntry         *gtk.Entry             `gtk:"hash_entry"`
	FileButton        *gtk.FileChooserButton `gtk:"file_button"`
	VerifyingSpinner  *gtk.Spinner           `gtk:"verifying_spinner"`
	ErrorDialog       *gtk.MessageDialog     `gtk:"error_dialog"`
	AboutDialog       *gtk.AboutDialog       `gtk:"about_dialog"`
	StatusStack       *gtk.Stack             `gtk:"status_stack"`
	StatusOkImage     *gtk.Image             `gtk:"status_ok_image"`
	StatusFailImage   *gtk.Image             `gtk:"status_fail_image"`
	HashLabel         *gtk.Label             `gtk:"hash_label"`
	AboutButton       *gtk.ModelButton       `gtk:"about_button"`

	signals map[string]interface{}
}

func New(appID, version string) *View {
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

		icon, err := gtk.ImageNewFromResource("/data/checksumo.svg")
		if err != nil {
			log.Fatal(err)
		}

		logo, err := icon.GetPixbuf().ScaleSimple(128, 128, gdk.INTERP_BILINEAR)
		if err != nil {
			log.Fatal(err)
		}

		view.AboutDialog.SetLogo(logo)
		view.AboutDialog.SetVersion(version)

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
