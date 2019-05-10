package timeutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var berlin, _ = time.LoadLocation("Europe/Berlin")
var sydney, _ = time.LoadLocation("Australia/Sydney")

func TestFixMissingTimezone(t *testing.T) {
	unzonedJan, err := time.Parse(time.RFC3339Nano, "2019-01-10T12:32:39.981423433Z")
	assert.NoError(t, err)
	unzonedJun, err := time.Parse(time.RFC3339Nano, "2019-06-10T12:32:39.981423433Z")
	assert.NoError(t, err)

	berlinJan := FixMissingTimezone(unzonedJan, berlin)
	berlinJun := FixMissingTimezone(unzonedJun, berlin)
	sydneyJan := FixMissingTimezone(unzonedJan, sydney)
	sydneyJun := FixMissingTimezone(unzonedJun, sydney)

	assert.Equal(t, time.Hour, unzonedJan.Sub(berlinJan))
	assert.Equal(t, 2 * time.Hour, unzonedJun.Sub(berlinJun))
	assert.Equal(t, 11* time.Hour, unzonedJan.Sub(sydneyJan))
	assert.Equal(t, 10* time.Hour, unzonedJun.Sub(sydneyJun))
}
