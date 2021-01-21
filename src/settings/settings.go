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
