package timeutil

import (
	"time"
)

const (
	UTC    = "UTC"
	LOCAL  = "Local"
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

// Takes a time that has wrongly been interpreted as UTC but was actually a local time value, and the location for
// that it should have been interpreted. Returns a timestamp with the error corrected. Considers the state of
// daylight saving time at the given location at that given point in time.
// This obviously must fail for one of the two duplicate hours during time adjustment, as that information is not
// present.
func FixMissingTimezone(unzoned time.Time, location *time.Location) time.Time {
	printed := unzoned.Format("2006-01-02T15:04:05.999999999Z")
	whitoutZone := printed[:len(printed)-1]
	zoned, _ := time.ParseInLocation("2006-01-02T15:04:05.999999999", whitoutZone, location)
	return zoned
}
