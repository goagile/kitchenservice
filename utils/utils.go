package utils

import (
	"time"
)

//
// DateTimeEq
//
func DateTimeEq(a, b time.Time) bool {
	return a.Day() == b.Day() &&
		a.Month() == b.Month() &&
		a.Year() == b.Year() &&
		a.Hour() == b.Hour() &&
		a.Minute() == b.Minute() &&
		a.Second() == b.Second()
}

//
// TestDateTime
//
var TestDateTime = time.Date(2020, time.October, 13, 23, 30, 10, 0, time.UTC)

//
// NewFakeClock
//
func NewFakeClock(dt time.Time) *fakeClock {
	return &fakeClock{dt}
}

type fakeClock struct {
	dt time.Time
}

func (fc *fakeClock) Now() time.Time {
	return fc.dt
}
