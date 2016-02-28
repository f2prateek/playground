package playground

import "time"

// Time returns the fixed current time in the go playground
// This time is equivalent to Tue, 10 Nov 2009 23:00:00 UTC.
func Time() time.Time {
	return time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
}
