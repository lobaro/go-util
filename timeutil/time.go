package timeutil

import (
	"time"

	"github.com/Sirupsen/logrus"
)

const (
	BERLIN = "Europe/Berlin"
)

func SetLocalTimezone(location string) {
	berlinLocation, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		logrus.WithError(err).Warn("Failed to load timezone for Europe/Berlin using system default.")
	} else {
		time.Local = berlinLocation
		logrus.Info("Loaded local timezone: Europe/Berlin")
	}
}
