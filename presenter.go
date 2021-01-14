package main

import (
	"context"
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

func (presenter *Presenter) SetFileOrHash() {
	// Update UI accordingly
}

func (presenter *Presenter) SetFile() {
	filePath := presenter.view.FileChooserButton.GetFilename()

	presenter.model.SetFilePath(filePath)
	hashType := presenter.model.DetectProvidedHashType()
	isGood := presenter.model.IsGoodToGo()

	presenter.view.MainHeaderBar.SetSubtitle(hashType)
	presenter.view.VerifyButton.SetSensitive(isGood)
	presenter.view.StatusStack.SetVisible(false)
}

func (presenter *Presenter) SetHash() {
	hashValue, _ := presenter.view.HashValueEntry.GetText()

	presenter.model.SetProvidedHash(hashValue)
	hashType := presenter.model.DetectProvidedHashType()
	isGood := presenter.model.IsGoodToGo()

	presenter.view.MainHeaderBar.SetSubtitle(hashType)
	presenter.view.VerifyButton.SetSensitive(isGood)
	presenter.view.StatusStack.SetVisible(false)
}

func (presenter *Presenter) StopHashing() {
	presenter.model.StopHashing()
}

func (presenter *Presenter) StartHashing() {
	presenter.model.CreateContext()

	presenter.view.ButtonStack.SetVisibleChild(presenter.view.CancelButton)
	presenter.view.StatusStack.SetVisibleChild(presenter.view.StatusSpinner)
	presenter.view.StatusStack.SetVisible(true)
	presenter.view.FileChooserButton.SetSensitive(false)
	presenter.view.HashValueEntry.SetSensitive(false)

	presenter.model.SetResultFunc(func(ok bool, err error) {
		if err == context.Canceled {
			presenter.view.StatusStack.SetVisible(false)
		} else if err != nil {
			presenter.view.ErrorDialog.FormatSecondaryText(err.Error())
			presenter.view.ErrorDialog.Run()
			presenter.view.ErrorDialog.Hide()
		} else if ok {
			presenter.view.StatusStack.SetVisibleChild(presenter.view.StatusOkImage)
		} else {
			presenter.view.StatusStack.SetVisibleChild(presenter.view.StatusFailImage)
		}

		presenter.view.ButtonStack.SetVisibleChild(presenter.view.VerifyButton)
		presenter.view.FileChooserButton.SetSensitive(true)
		presenter.view.HashValueEntry.SetSensitive(true)
	})

	go presenter.model.StartHashing()
}
