package ticket

import (
	"time"
)

//
// Clock
//
type Clock interface {
	Now() time.Time
}

//
// DefaultClock
//
var DefaultClock Clock = NewSystemClock()

//
// NewSystemClock
//
func NewSystemClock() *sysclock {
	return &sysclock{}
}

type sysclock struct {}

//
// Now
//
func (c *sysclock) Now() time.Time  {
	return time.Now()
}
