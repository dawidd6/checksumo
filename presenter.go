package main

import (
	"context"

	"github.com/gotk3/gotk3/glib"
)

type Presenter struct {
	view  *View
	model *Model
}

func NewPresenter(view *View) *Presenter {
	return &Presenter{
		view:  view,
		model: NewModel(),
	}
}

func (presenter *Presenter) SetFile() {
	presenter.model.filePath = presenter.view.FileChooserButton.GetFilename()

	hashType := presenter.model.DetectProvidedHashType()
	isGood := presenter.model.IsGoodToGo()

	presenter.view.MainHeaderBar.SetSubtitle(hashType)
	presenter.view.VerifyButton.SetSensitive(isGood)
	presenter.view.HashValueEntry.SetProgressFraction(0.0)
}

func (presenter *Presenter) SetHash() {
	presenter.model.providedHash, _ = presenter.view.HashValueEntry.GetText()

	hashType := presenter.model.DetectProvidedHashType()
	isGood := presenter.model.IsGoodToGo()

	presenter.view.MainHeaderBar.SetSubtitle(hashType)
	presenter.view.VerifyButton.SetSensitive(isGood)
	presenter.view.HashValueEntry.SetProgressFraction(0.0)
}

func (presenter *Presenter) StopHashing() {
	presenter.model.StopHashing()
}

func (presenter *Presenter) StartHashing() {
	presenter.model.CreateContext()

	presenter.view.ButtonStack.SetVisibleChild(presenter.view.CancelButton)
	presenter.view.FileChooserButton.SetSensitive(false)
	presenter.view.HashValueEntry.SetSensitive(false)

	presenter.model.resultFunc = func(ok bool, err error) {
		glib.IdleAdd(func() {
			if err == context.Canceled {
			} else if err != nil {
				presenter.view.ErrorDialog.FormatSecondaryText(err.Error())
				presenter.view.ErrorDialog.Run()
				presenter.view.ErrorDialog.Hide()
			} else if ok {
				presenter.view.ResultOkDialog.Run()
				presenter.view.ResultOkDialog.Hide()
			} else {
				presenter.view.ResultFailDialog.Run()
				presenter.view.ResultFailDialog.Hide()
			}

			presenter.view.ButtonStack.SetVisibleChild(presenter.view.VerifyButton)
			presenter.view.FileChooserButton.SetSensitive(true)
			presenter.view.HashValueEntry.SetSensitive(true)
		})
	}

	presenter.model.progressFunc = func(progressFraction float32) {
		glib.IdleAdd(presenter.view.HashValueEntry.SetProgressFraction, progressFraction)
	}

	go presenter.model.StartHashing()
}
