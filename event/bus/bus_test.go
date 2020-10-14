package bus

import (
	"testing"

	"github.com/goagile/kitchenservice/event"
)

//
// Add
//
func Test_Bus_Add(t *testing.T) {
	want := 1
	bus := New()
	bus.Add(&funcHandler{})

	got := len(bus.handlers)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// AddFunc
//
func Test_Bus_AddFunc(t *testing.T) {
	want := 1
	bus := New()
	bus.AddFunc(func(e event.Event) {})

	got := len(bus.handlers)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// Publish
//
func Test_Bus_Publish(t *testing.T) {
	want := true
	bus := New()

	got := false
	bus.AddFunc(func(e event.Event) {
		got = true
	})

	bus.Publish(new(event.Event))

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}
