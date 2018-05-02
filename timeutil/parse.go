package timeutil

import (
	"time"
	"strconv"
)

var timeFormats = []string{
	time.RFC3339Nano,
	time.RFC3339,
	"2006-01-02T15:04:05.999999999 07:00", // time.RFC3339Nano with space insted of Z
	"2006-01-02T15:04:05 07:00",           // time.RFC3339 with space insted of Z
	"2006-01-02T15:04:05.999999999",
	"2006-01-02T15:04:05",
}

func AddTimeFormat(format string) {
	timeFormats = append(timeFormats, format)
}

func GetTimeFormats() []string {
	return timeFormats[:]
}

func SetTimeFormats(formats []string) {
	timeFormats = formats[:]
}

func ParseTime(s string) time.Time {
	return ParseTimeWith(s, timeFormats...)
}

func MustParse(layout string, value string) time.Time {
	if t, err := time.Parse(layout, value); err != nil {
		panic(err)
	} else {
		return t
	}
}

func ParseTimeWith(s string, formats... string) time.Time {
	// First try to parse as int
	i, err := strconv.ParseInt(s, 10, 64)
	t := time.Unix(0, i)

	if err == nil {
		// We use a heuristic to detect if it's nanos or seconds
		if i < 32503680000 { // max value: 01.01.3000 12:00:00, after interpret as millis
			//logrus.Info("Seconds: ", i)
			return time.Unix(i, 0)
		} else { // min value: 12.1.1971, 05:48:00, before interpret as seconds
			//logrus.Info("Milliseconds: ", i)
			return time.Unix(0, i*1e6)
		}
	}

	//logrus.Info("Parsing time: ", s)
	for _, f := range formats {
		t, err = time.Parse(f, s)
		if err == nil {
			return t
		}
	}

	return t
}
