package controller

import (
	"context"

	"github.com/gotk3/gotk3/glib"

	"github.com/dawidd6/checksumo/pkg/model"
	"github.com/dawidd6/checksumo/pkg/view"
)

type Controller struct {
	View  *view.View
	Model *model.Model

	ctx    context.Context
	cancel context.CancelFunc
}

func New(v *view.View, m *model.Model) *Controller {
	// Construct controller
	controller := &Controller{
		View:  v,
		Model: m,
	}

	// Connect handlers to signals when widgets are ready
	controller.View.Application.Connect("ready", func() {
		controller.View.FileChooserButton.Connect("file-set", controller.onFileButtonFileSet)
		controller.View.HashValueEntry.Connect("changed", controller.onHashEntryChanged)
		controller.View.HashValueEntry.Connect("activate", controller.onHashEntryActivate)
		controller.View.VerifyButton.Connect("clicked", controller.onVerifyButtonClicked)
		controller.View.CancelButton.Connect("clicked", controller.onCancelButtonClicked)
		controller.View.SettingsButton.Connect("clicked", controller.onSettingsButtonClicked)
		controller.View.SaveButton.Connect("clicked", controller.onSaveButtonClicked)
	})

	return controller
}

func (controller *Controller) onFileButtonFileSet() {
	go func() {
		// Get inputs
		filePath := controller.View.FileChooserButton.GetFilename()
		hashValue, _ := controller.View.HashValueEntry.GetText()
		// Detect hash type
		hashType := controller.Model.DetectType(hashValue)

		// Check conditions and decide if we can allow verification yet
		allowVerify := false
		if hashType != "" && hashValue != "" && filePath != "" {
			allowVerify = true
		}

		// Update UI accordingly
		glib.IdleAdd(func() {
			controller.View.MainHeaderBar.SetSubtitle(hashType)
			controller.View.VerifyButton.SetSensitive(allowVerify)
			controller.View.StatusStack.SetVisible(false)
		})
	}()
}

func (controller *Controller) onHashEntryChanged() {
	// Perform the same checks
	controller.onFileButtonFileSet()
}

func (controller *Controller) onHashEntryActivate() {
	// Treat ENTER button as a confirmation
	controller.onVerifyButtonClicked()
}

func (controller *Controller) onCancelButtonClicked() {
	// Cancel context and stop verification
	controller.cancel()
}

func (controller *Controller) onVerifyButtonClicked() {
	go func() {
		var err error

		// Create context
		controller.ctx, controller.cancel = context.WithCancel(context.Background())
		defer controller.cancel()

		// Initial UI updates when verification starts
		glib.IdleAdd(func() {
			controller.View.ButtonStack.SetVisibleChild(controller.View.CancelButton)
			controller.View.StatusStack.SetVisibleChild(controller.View.StatusSpinner)
			controller.View.StatusStack.SetVisible(true)
			controller.View.FileChooserButton.SetSensitive(false)
			controller.View.HashValueEntry.SetSensitive(false)
		})

		// Update UI on return
		defer glib.IdleAdd(func() {
			controller.View.ButtonStack.SetVisibleChild(controller.View.VerifyButton)
			controller.View.FileChooserButton.SetSensitive(true)
			controller.View.HashValueEntry.SetSensitive(true)

			switch err {
			case nil:
				return
			case context.Canceled:
				// User cancelled operation
				controller.View.StatusStack.SetVisible(false)
			default:
				// Display error dialog
				controller.View.ErrorDialog.FormatSecondaryText(err.Error())
				controller.View.ErrorDialog.Run()
				controller.View.ErrorDialog.Hide()
			}
		})

		// Get user inputs
		filePath := controller.View.FileChooserButton.GetFilename()
		hashValueProvided, _ := controller.View.HashValueEntry.GetText()
		if err != nil {
			return
		}

		// Compute file hash
		hashValueComputed, err := controller.Model.ComputeHash(controller.ctx, filePath)
		if err != nil {
			return
		}

		// Determine result image
		resultImage := controller.View.StatusFailImage
		if hashValueProvided == hashValueComputed {
			resultImage = controller.View.StatusOkImage
		}

		// Update UI to signalize hash comparison to user
		glib.IdleAdd(func() {
			controller.View.StatusStack.SetVisibleChild(resultImage)
		})
	}()
}

func (controller *Controller) onSettingsButtonClicked() {
	controller.View.SettingsWindow.Present()
}

func (controller *Controller) onSaveButtonClicked() {
	controller.View.SettingsWindow.Hide()
}
