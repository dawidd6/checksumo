package presenters

import (
	"context"

	"github.com/dawidd6/checksumo/src/models"
	"github.com/gotk3/gotk3/glib"
)

type MainView interface {
	GetFile() string
	GetHash() string
	OnProcessStart()
	OnProcessStop()
	OnResultSuccess()
	OnResultFailure()
	OnResultError(error)
	OnProgressUpdate(float64)
	OnFileOrHashSet(bool, string)
}

type MainPresenter struct {
	view  MainView
	model *models.MainModel
}

func NewMainPresenter(view MainView) *MainPresenter {
	return &MainPresenter{
		view:  view,
		model: models.NewMainModel(),
	}
}

func (presenter *MainPresenter) setFileOrHash() {
	// Detect hash type from provided hash
	hashType := presenter.model.DetectType()
	// Check if every needed information is provided
	isReady := presenter.model.IsReady()
	// Display detected hash and unblock starting button if ready
	presenter.view.OnFileOrHashSet(isReady, hashType)
}

func (presenter *MainPresenter) SetFile() {
	// Get file from view and set it in model
	filePath := presenter.view.GetFile()
	presenter.model.SetFile(filePath)
	// Run common function
	presenter.setFileOrHash()
}

func (presenter *MainPresenter) SetHash() {
	// Get hash from view and set it in model
	hashValue := presenter.view.GetHash()
	presenter.model.SetHash(hashValue)
	// Run common function
	presenter.setFileOrHash()
}

func (presenter *MainPresenter) StopHashing() {
	// This essentially comes down to calling context.CancelFunc
	presenter.model.StopHashing()
}

func (presenter *MainPresenter) StartHashing() {
	// Check if ready
	if !presenter.model.IsReady() {
		return
	}

	// Prepare hashing by creating context with cancel
	presenter.model.PrepareHashing()

	// Prepare view
	presenter.view.OnProcessStart()

	// Keep updating the progress in view
	progressSource, _ := glib.TimeoutAdd(10, func() bool {
		presenter.view.OnProgressUpdate(presenter.model.GetProgress())
		return true
	})

	// Start separate goroutine for hashing process as it can take a long time
	// and we don't want to freeze the UI
	go func() {
		ok, err := presenter.model.StartHashing()

		// Stop updating the progress in view
		glib.SourceRemove(progressSource)

		// Finalize the progress visualization
		glib.IdleAdd(presenter.view.OnProgressUpdate, presenter.model.GetProgress())

		// Show the appropriate result
		if err == context.Canceled {
			// NOOP
		} else if err != nil {
			glib.IdleAdd(presenter.view.OnResultError, err)
		} else if ok {
			glib.IdleAdd(presenter.view.OnResultSuccess)
		} else {
			glib.IdleAdd(presenter.view.OnResultFailure)
		}

		// Return the view to the state before hashing process
		glib.IdleAdd(presenter.view.OnProcessStop)
	}()
}
