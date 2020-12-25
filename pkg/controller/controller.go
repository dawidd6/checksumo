package controller

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"

	"github.com/dawidd6/ghashverifier/pkg/model"
	"github.com/dawidd6/ghashverifier/pkg/view"
)

type Controller struct {
	View  *view.View
	Model *model.Model
}

func New(v *view.View, m *model.Model) *Controller {
	controller := &Controller{
		View:  v,
		Model: m,
	}

	signals := map[string]interface{}{
		"on_file_button_file_set":    controller.onFileButtonFileSet,
		"on_hash_entry_changed":      controller.onHashEntryChanged,
		"on_verify_button_clicked":   controller.onVerifyButtonClicked,
		"on_settings_button_clicked": controller.onSettingsButtonClicked,
	}

	v.SetSignals(signals)

	return controller
}

func (controller *Controller) Run() {
	os.Exit(controller.View.Application.Run(os.Args))
}

func (controller *Controller) onFileButtonFileSet() {
	controller.View.HashLabel.SetText("")
	controller.View.StatusStack.SetVisibleChild(controller.View.VerifyButton)
}

func (controller *Controller) onHashEntryChanged() {
	controller.View.HashLabel.SetText("")
	controller.View.StatusStack.SetVisibleChild(controller.View.VerifyButton)
}

func (controller *Controller) onVerifyButtonClicked() {
	go func() {
		var err error

		// Show spinner
		glib.IdleAdd(controller.View.StatusStack.SetVisibleChild, controller.View.VerifyingSpinner)

		// Cleanup on return
		defer glib.IdleAdd(func() {
			if err != nil {
				controller.View.HashLabel.SetText("")
				controller.View.StatusStack.SetVisibleChild(controller.View.VerifyButton)
				controller.View.ErrorDialog.FormatSecondaryText(err.Error())
				controller.View.ErrorDialog.Run()
				controller.View.ErrorDialog.Hide()
			}
		})

		// Get user inputs
		filePath := controller.View.FileButton.GetFilename()
		providedHash, err := controller.View.HashEntry.GetText()
		if err != nil {
			return
		}

		// Detect hash type
		hashType, err := controller.Model.DetectType(providedHash)
		if err != nil {
			return
		}

		// Show detected hash type
		glib.IdleAdd(controller.View.HashLabel.SetText, hashType)

		// Compute file hash
		gotHash, err := controller.Model.ComputeHash(filePath)
		if err != nil {
			return
		}

		// Signalize hash comparison to user
		if providedHash == gotHash {
			glib.IdleAdd(controller.View.StatusStack.SetVisibleChild, controller.View.StatusOkImage)
		} else {
			glib.IdleAdd(controller.View.StatusStack.SetVisibleChild, controller.View.StatusFailImage)
		}
	}()
}

func (controller *Controller) onSettingsButtonClicked() {
	log.Println("settings click")
}
