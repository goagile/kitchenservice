package utils

import (
	"testing"
	"time"
)

//
// DateTimeEq
//

func Test_DateTimeEq_True(t *testing.T) {
	want := true
	a := November(12, 2019, 23, 10)
	b := November(12, 2019, 23, 10)

	got := DateTimeEq(a, b)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// November
//
func November(d int, y int, hh int, mm int) time.Time {
	return time.Date(y, time.November, d, hh, mm, 0, 0, time.UTC)
}
