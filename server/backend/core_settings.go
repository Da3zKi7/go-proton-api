package backend

import (
	"github.com/Da3zKi7/go-proton-api"
)

func newUserSettings() proton.UserSettings {
	return proton.UserSettings{Telemetry: proton.SettingEnabled, CrashReports: proton.SettingEnabled}
}
