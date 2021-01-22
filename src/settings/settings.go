package settings

import (
	"github.com/dawidd6/checksumo/src/constants"
	"github.com/gotk3/gotk3/glib"
)

var settings = glib.SettingsNew(constants.AppID)

func ShowNotifications(value ...bool) bool {
	key := "show-notifications"
	if value == nil {
		return settings.GetBoolean(key)
	}
	return settings.SetBoolean(key, value[0])
}

func RememberWindowPosition(value ...bool) bool {
	key := "remember-window-position"
	if value == nil {
		return settings.GetBoolean(key)
	}
	return settings.SetBoolean(key, value[0])
}

func RememberWindowSize(value ...bool) bool {
	key := "remember-window-size"
	if value == nil {
		return settings.GetBoolean(key)
	}
	return settings.SetBoolean(key, value[0])
}

func RememberDirectory(value ...bool) bool {
	key := "remember-directory"
	if value == nil {
		return settings.GetBoolean(key)
	}
	return settings.SetBoolean(key, value[0])
}

func SavedDirectory(value ...string) string {
	key := "saved-directory"
	if value == nil {
		return settings.GetString(key)
	}
	settings.SetString(key, value[0])
	return ""
}

func SavedWindowWidth(value ...int) int {
	key := "saved-window-width"
	if value == nil {
		return settings.GetInt(key)
	}
	settings.SetInt(key, value[0])
	return -1
}

func SavedWindowHeight(value ...int) int {
	key := "saved-window-height"
	if value == nil {
		return settings.GetInt(key)
	}
	settings.SetInt(key, value[0])
	return -1
}

func SavedWindowPositionX(value ...int) int {
	key := "saved-window-position-x"
	if value == nil {
		return settings.GetInt(key)
	}
	settings.SetInt(key, value[0])
	return -1
}

func SavedWindowPositionY(value ...int) int {
	key := "saved-window-position-y"
	if value == nil {
		return settings.GetInt(key)
	}
	settings.SetInt(key, value[0])
	return -1
}
