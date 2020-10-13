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
// Clock
//
type Clock interface {
	Now() time.Time
}

//
// NewSystemClock
//
func NewSystemClock() *sysclock {
	return &sysclock{}
}

type sysclock struct{}

//
// Now
//
func (c *sysclock) Now() time.Time {
	return time.Now()
}

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

//
// TestDateTime
//
var TestDateTime = time.Date(
	2020, time.October, 13,
	23, 30, 10, 0,
	time.UTC,
)
