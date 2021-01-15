package controller

import (
	"os"

	"github.com/dawidd6/checksumo/src/model"
	"github.com/dawidd6/checksumo/src/view"

	"github.com/gotk3/gotk3/glib"
)

type Controller struct {
	v *view.View
	m *model.Model
}

func New(v *view.View, m *model.Model) *Controller {
	controller := &Controller{
		v: v,
		m: m,
	}

	v.Application.ConnectAfter("activate", func() {
		v.FileChooserButton.Connect("file-set", controller.SetFile)
		v.HashValueEntry.Connect("changed", controller.SetHash)
		v.HashValueEntry.Connect("activate", controller.StartHashing)
		v.VerifyButton.Connect("clicked", controller.StartHashing)
		v.CancelButton.Connect("clicked", controller.StopHashing)
	})

	return controller
}

func (controller *Controller) Run() int {
	return controller.v.Application.Run(os.Args)
}

func (controller *Controller) SetFile() {
	filePath := controller.v.FileChooserButton.GetFilename()
	controller.m.SetFile(filePath)

	hashType := controller.m.DetectType()
	isReady := controller.m.IsReady()

	if controller.v.MainHeaderBar.GetSubtitle() != hashType {
		controller.v.MainHeaderBar.SetSubtitle(hashType)
	}
	if controller.v.VerifyButton.GetSensitive() != isReady {
		controller.v.VerifyButton.SetSensitive(!isReady)
	}
	if controller.v.HashValueEntry.GetProgressFraction() > 0 {
		controller.v.HashValueEntry.SetProgressFraction(0.0)
	}
}

func (controller *Controller) SetHash() {
	hashValue, _ := controller.v.HashValueEntry.GetText()
	controller.m.SetHash(hashValue)

	hashType := controller.m.DetectType()
	isReady := controller.m.IsReady()

	if controller.v.MainHeaderBar.GetSubtitle() != hashType {
		controller.v.MainHeaderBar.SetSubtitle(hashType)
	}
	if controller.v.VerifyButton.GetSensitive() != isReady {
		controller.v.VerifyButton.SetSensitive(isReady)
	}
	if controller.v.HashValueEntry.GetProgressFraction() > 0 {
		controller.v.HashValueEntry.SetProgressFraction(0.0)
	}
}

func (controller *Controller) StopHashing() {
	controller.m.StopHashing()
}

func (controller *Controller) StartHashing() {
	if !controller.m.IsReady() {
		return
	}

	controller.m.PrepareHashing()

	controller.v.ButtonStack.SetVisibleChild(controller.v.CancelButton)
	controller.v.FileChooserButton.SetSensitive(false)
	controller.v.HashValueEntry.SetSensitive(false)

	glib.TimeoutAdd(10, func() bool {
		controller.v.HashValueEntry.SetProgressFraction(controller.m.GetProgress())
		return controller.m.IsBusy()
	})

	controller.m.SetResultFunc(func(ok bool, err error) {
		glib.IdleAdd(func() {
			if err != nil {
				controller.v.ErrorDialog.FormatSecondaryText(err.Error())
				controller.v.ErrorDialog.Run()
				controller.v.ErrorDialog.Hide()
			} else if ok {
				controller.v.ResultOkDialog.Run()
				controller.v.ResultOkDialog.Hide()
			} else {
				controller.v.ResultFailDialog.Run()
				controller.v.ResultFailDialog.Hide()
			}

			controller.v.ButtonStack.SetVisibleChild(controller.v.VerifyButton)
			controller.v.FileChooserButton.SetSensitive(true)
			controller.v.HashValueEntry.SetSensitive(true)
		})
	})

	go controller.m.StartHashing()
}
