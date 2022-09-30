package helper

import (
	"time"
)

type timeComparison struct {
	diff time.Duration // Time difference in minutes
	isGt bool          // Is time.Now() greater than ttl
}

// CompareTimeNow To Compare between two dates: time.Now() and date to compare formatted in UnixMilli()
func CompareTimeNow(ttl int64) *timeComparison {
	var isGt bool
	t := time.Now().UTC()
	timeParse := time.UnixMilli(ttl).UTC()

	if t.After(timeParse) {
		isGt = true
	} else {
		isGt = false
	}

	return &timeComparison{
		diff: timeParse.Sub(t),
		isGt: isGt,
	}
}
