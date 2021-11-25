package util

import "time"

const TIME_ZONE = "Asia/Tokyo"

var loc *time.Location

func init() {
	l, err := time.LoadLocation(TIME_ZONE)
	if err != nil {
		panic(err)
	}
	loc = l
}

// LocalNow returns the current time.
// The time zone is defined in this package.
func LocalNow() time.Time {
	return time.Now().In(loc)
}
