package timeutil

import "testing"

func TestParseTime(t *testing.T) {

	// Parse Milli seconds
	time := ParseTime("1523456789101").UTC()

	if time.String() != "2018-04-11 14:26:29.101 +0000 UTC" {
		t.Error("Time: ", time.String())
	}

}
