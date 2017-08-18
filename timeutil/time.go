package timeutil

import (
	"time"
)

const (
	BERLIN = "Europe/Berlin"
)

func SetLocalTimezone(location string) error {
	berlinLocation, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		return err
	}
	time.Local = berlinLocation
	return nil
}
