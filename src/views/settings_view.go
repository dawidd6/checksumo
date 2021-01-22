package views

import (
	"github.com/dawidd6/checksumo/src/settings"
	"github.com/dawidd6/checksumo/src/utils"

	"github.com/gotk3/gotk3/gtk"
)

type settingsView struct {
	SettingsWindow              *gtk.Window
	SettingsHeaderBar           *gtk.HeaderBar
	SaveButton                  *gtk.Button
	ShowNotificationsCheck      *gtk.CheckButton
	RememberDirectoryCheck      *gtk.CheckButton
	RememberWindowSizeCheck     *gtk.CheckButton
	RememberWindowPositionCheck *gtk.CheckButton
}

func NewSettingsView() *settingsView {
	return &settingsView{}
}

func (view *settingsView) Activate() {
	// Bind widgets
	utils.BindWidgets(view, "/com/github/dawidd6/checksumo/ui/settings_view.ui")

	// Display current settings state
	view.ShowNotificationsCheck.SetActive(settings.ShowNotifications())
	view.RememberDirectoryCheck.SetActive(settings.RememberDirectory())
	view.RememberWindowSizeCheck.SetActive(settings.RememberWindowSize())
	view.RememberWindowPositionCheck.SetActive(settings.RememberWindowPosition())

	// Connect handlers to events
	view.SaveButton.Connect("clicked", view.SettingsWindow.Close)
	view.ShowNotificationsCheck.Connect("toggled", func() {
		settings.ShowNotifications(view.ShowNotificationsCheck.GetActive())
	})
	view.RememberDirectoryCheck.Connect("toggled", func() {
		settings.RememberDirectory(view.RememberDirectoryCheck.GetActive())
	})
	view.RememberWindowSizeCheck.Connect("toggled", func() {
		settings.RememberWindowSize(view.RememberWindowSizeCheck.GetActive())
	})
	view.RememberWindowPositionCheck.Connect("toggled", func() {
		settings.RememberWindowPosition(view.RememberWindowPositionCheck.GetActive())
	})

	// Show settings window
	view.SettingsWindow.Present()
}
