package ticket

import (
	"testing"
	"time"
)

//
// Created
//
func Test_NewTicket_Has_Created_State(t *testing.T) {
	want := Created
	tic := New(TicketID(1))

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_NewTicket_CreatedAt(t *testing.T) {
	DefaultClock = newFakeClock(TestDateTime)
	want := TestDateTime
	tic := New(TicketID(1))

	got := tic.CreatedAt

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// Cancelled
//
func Test_NewTicket_May_Be_Cancelled(t *testing.T) {
	want := Cancelled
	tic := New(TicketID(1))
	tic.Cancel()

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_CancelFromAcceptedIsNotValid(t *testing.T) {
	tic := New(TicketID(1))
	tic.Accept()

	err := tic.Cancel()

	if err != CancelFromAcceptedIsNotValid {
		t.Fatalf("\nshould return error")
	}
}

func Test_CancelFromPreparedIsNotValid(t *testing.T) {
	tic := New(TicketID(1))
	tic.Accept()
	tic.Prepare()

	err := tic.Cancel()

	if err != CancelFromPreparedIsNotValid {
		t.Fatalf("\nshould return error")
	}
}

func Test_PreparedTicket_May_Be_Cancelled(t *testing.T) {
	want := Cancelled
	tic := New(TicketID(1))
	tic.Accept()
	tic.Prepare()
	tic.ReadyToPickUp()
	tic.Cancel()

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_CancelledAt(t *testing.T) {
	DefaultClock = newFakeClock(TestDateTime)
	want := TestDateTime
	tic := New(TicketID(1))
	tic.Cancel()

	got := tic.CancelledAt

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// Accepted
//
func Test_New_Accepted_Ticket_Has_Accepted_State(t *testing.T) {
	want := Accepted
	tic := New(TicketID(1))
	tic.Accept()

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_Accept_Twice_Is_Ok(t *testing.T) {
	want := Accepted
	tic := New(TicketID(1))
	tic.Accept()
	tic.Accept()

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_AcceptedFromPreparedIsNotValid(t *testing.T) {
	tic := New(TicketID(1))
	tic.Accept()
	tic.Prepare()

	err := tic.Accept()

	if err != AcceptedFromPreparedIsNotValid {
		t.Fatalf("\nshould return error")
	}
}

func Test_AcceptedFromReadyToPickUpIsNotValid(t *testing.T) {
	tic := New(TicketID(1))
	tic.Accept()
	tic.Prepare()
	tic.ReadyToPickUp()

	err := tic.Accept()

	if err != AcceptedFromReadyToPickUpIsNotValid {
		t.Fatalf("\nshould return error")
	}
}

func Test_AcceptedAt(t *testing.T) {
	DefaultClock = newFakeClock(TestDateTime)
	want := TestDateTime
	tic := New(TicketID(1))
	tic.Accept()

	got := tic.AcceptedAt

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// Prepared
//
func Test_Prepare_Accepted_Ticket_Set_Prepared_State(t *testing.T) {
	want := Prepared
	tic := New(TicketID(1))
	tic.Accept()
	tic.Prepare()

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_Prepare_Twice_Is_Ok(t *testing.T) {
	want := Prepared
	tic := New(TicketID(1))
	tic.Accept()
	tic.Prepare()
	tic.Prepare()

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_Created_To_Prepared_Is_Not_Valid(t *testing.T) {
	tic := New(TicketID(1))

	err := tic.Prepare()

	if err != PrepareFromCreatedIsNotValid {
		t.Fatalf("\nshould return error")
	}
}

func Test_PrepareFromReadyToPickUpIsNotValid(t *testing.T) {
	tic := New(TicketID(1))
	tic.Accept()
	tic.Prepare()
	tic.ReadyToPickUp()

	err := tic.Prepare()

	if err != PrepareFromReadyToPickUpIsNotValid {
		t.Fatalf("\nshould return error")
	}
}

func Test_PreparedAt(t *testing.T) {
	DefaultClock = newFakeClock(TestDateTime)
	want := TestDateTime
	tic := New(TicketID(1))
	tic.Accept()
	tic.Prepare()

	got := tic.PreparedAt

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// ReadyToPickUp
//
func Test_Prepare_to_ReadyToPickUp_Set_ReadyToPickUp_State(t *testing.T) {
	want := ReadyToPickUp
	tic := New(TicketID(1))
	tic.Accept()
	tic.Prepare()
	tic.ReadyToPickUp()

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_ReadyToPickUp_Twice_Is_Ok(t *testing.T) {
	want := ReadyToPickUp
	tic := New(TicketID(1))
	tic.Accept()
	tic.Prepare()
	tic.ReadyToPickUp()
	tic.ReadyToPickUp()

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_ReadyToPickUpFromCreatedIsNotValid(t *testing.T) {
	tic := New(TicketID(1))

	err := tic.ReadyToPickUp()

	if err != ReadyToPickUpFromCreatedIsNotValid {
		t.Fatalf("\nshould return error")
	}
}

func Test_ReadyToPickUpFromAcceptedIsNotValid(t *testing.T) {
	tic := New(TicketID(1))
	tic.Accept()

	err := tic.ReadyToPickUp()

	if err != ReadyToPickUpFromAcceptedIsNotValid {
		t.Fatalf("\nshould return error")
	}
}

func Test_ReadyForPickUpAt(t *testing.T) {
	DefaultClock = newFakeClock(TestDateTime)
	want := TestDateTime
	tic := New(TicketID(1))
	tic.Accept()
	tic.Prepare()
	tic.ReadyToPickUp()

	got := tic.ReadyForPickUpAt

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// newFakeClock
//
func newFakeClock(dt time.Time) *fakeClock {
	return &fakeClock{dt}
}

var TestDateTime = time.Date(2020, time.October, 13, 23, 30, 10, 0, time.UTC)

type fakeClock struct {
	dt time.Time
}

//
// Now
//
func (fc *fakeClock) Now() time.Time {
	return fc.dt
}
