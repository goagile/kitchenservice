package event

import (
	"testing"
)

//
// AddHandler
//
func Test_Bus_AddHandler(t *testing.T) {
	want := 1
	bus := &Bus{}
	bus.AddHandler(&funcHandler{})

	got := len(bus.handlers)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// AddFuncHandler
//
func Test_Bus_AddFuncHandler(t *testing.T) {
	want := 1
	bus := &Bus{}
	bus.AddFuncHandler(func(e Event) {})

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
	bus := &Bus{}

	got := false
	bus.AddFuncHandler(func(e Event) {
		got = true
	})

	bus.Publish(EmptyEvent{})

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}
