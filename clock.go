package golog

import (
	"time"
)

// Clock holds the string formats of time to be used inside the log and for the logfile name
type Clock struct {
	timeFormat string
	dateFormat string
}

func (c *Clock) nowTime() string {
	return time.Now().Format(c.timeFormat)
}

func (c *Clock) nowDate() string {
	return time.Now().Format(c.dateFormat)
}

// NewClock returns a clock which uses the specified format
// [time.Format()] formatting is used
func NewClock(timeFormat string, dateFormat string) *Clock {
	return &Clock{timeFormat, dateFormat}
}

// NewClock returns a clock which a default format
func NewDefaultClock() *Clock {
	return &Clock{
		timeFormat: "15:04:05.000",
		dateFormat: time.RFC1123,
	}
}
