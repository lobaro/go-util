package timeutil

import (
	"time"
)

const (
	BERLIN = "Europe/Berlin"
)

const (
	FORMAT_DATE_DE         = "02.01.2006"
	FORMAT_TIME_DE         = "15:04:05"
	FORMAT_TZ              = "-0700 MST"
	FORMAT_TIME_DE_TZ      = FORMAT_TIME_DE + " " + FORMAT_TZ
	FORMAT_DATE_TIME_DE    = FORMAT_DATE_DE + " " + FORMAT_TIME_DE
	FORMAT_DATE_TIME_DE_TZ = FORMAT_DATE_DE + " " + FORMAT_TIME_DE + " " + FORMAT_TZ
)

func SetLocalTimezone(location string) error {
	berlinLocation, err := time.LoadLocation(location)
	if err != nil {
		return err
	}
	time.Local = berlinLocation
	return nil
}
