package main

import (
	"context"
	"time"

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

	if presenter.view.MainHeaderBar.GetSubtitle() != hashType {
		presenter.view.MainHeaderBar.SetSubtitle(hashType)
	}
	if presenter.view.VerifyButton.GetSensitive() != isGood {
		presenter.view.VerifyButton.SetSensitive(isGood)
	}
	if presenter.view.HashValueEntry.GetProgressFraction() > 0 {
		presenter.view.HashValueEntry.SetProgressFraction(0.0)
	}
}

func (presenter *Presenter) SetHash() {
	presenter.model.providedHash, _ = presenter.view.HashValueEntry.GetText()

	hashType := presenter.model.DetectProvidedHashType()
	isGood := presenter.model.IsGoodToGo()

	if presenter.view.MainHeaderBar.GetSubtitle() != hashType {
		presenter.view.MainHeaderBar.SetSubtitle(hashType)
	}
	if presenter.view.VerifyButton.GetSensitive() != isGood {
		presenter.view.VerifyButton.SetSensitive(isGood)
	}
	if presenter.view.HashValueEntry.GetProgressFraction() > 0 {
		presenter.view.HashValueEntry.SetProgressFraction(0.0)
	}
}

func (presenter *Presenter) StopHashing() {
	presenter.model.StopHashing()
}

func (presenter *Presenter) StartHashing() {
	track := true

	if !presenter.model.IsGoodToGo() {
		return
	}

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

		track = false
	}

	go func() {
		for track {
			time.Sleep(time.Millisecond * 10)
			progress := float32(presenter.model.currentSize) / float32(presenter.model.totalSize)
			glib.IdleAdd(presenter.view.HashValueEntry.SetProgressFraction, progress)
		}
	}()

	go presenter.model.StartHashing()
}
