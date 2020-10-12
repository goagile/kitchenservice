package ticket

import (
	"testing"
)

//
// Created
//
func Test_NewTicket_Has_Created_State(t *testing.T) {
	want := Created
	tic := NewTicket()

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// Cancelled
//
func Test_NewTicket_May_Be_Cancelled(t *testing.T) {
	want := Cancelled
	tic := NewTicket()
	tic.Cancel()

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_CancelFromAcceptedIsNotValid(t *testing.T) {
	tic := NewTicket()
	tic.Accept()
	
	err := tic.Cancel()
	
	if err != CancelFromAcceptedIsNotValid {
		t.Fatalf("\nshould return error")
	}
}

func Test_CancelFromPreparedIsNotValid(t *testing.T) {
	tic := NewTicket()
	tic.Accept()
	tic.Prepare()
	
	err := tic.Cancel()
	
	if err != CancelFromPreparedIsNotValid {
		t.Fatalf("\nshould return error")
	}
}

func Test_PreparedTicket_May_Be_Cancelled(t *testing.T) {
	want := Cancelled
	tic := NewTicket()
	tic.Accept()
	tic.Prepare()
	tic.ReadyToPickUp()
	tic.Cancel()

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

//
// Accepted
//
func Test_New_Accepted_Ticket_Has_Accepted_State(t *testing.T) {
	want := Accepted
	tic := NewTicket()
	tic.Accept()

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_Accept_Twice_Is_Ok(t *testing.T) {
	want := Accepted
	tic := NewTicket()
	tic.Accept()
	tic.Accept()

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_AcceptedFromPreparedIsNotValid(t *testing.T) {
	tic := NewTicket()
	tic.Accept()
	tic.Prepare()
	
	err := tic.Accept()
	
	if err != AcceptedFromPreparedIsNotValid {
		t.Fatalf("\nshould return error")
	}
}

func Test_AcceptedFromReadyToPickUpIsNotValid(t *testing.T) {
	tic := NewTicket()
	tic.Accept()
	tic.Prepare()
	tic.ReadyToPickUp()
	
	err := tic.Accept()
	
	if err != AcceptedFromReadyToPickUpIsNotValid {
		t.Fatalf("\nshould return error")
	}
}

//
// Prepared
//
func Test_Prepare_Accepted_Ticket_Set_Prepared_State(t *testing.T) {
	want := Prepared
	tic := NewTicket()
	tic.Accept()
	tic.Prepare()

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_Prepare_Twice_Is_Ok(t *testing.T) {
	want := Prepared
	tic := NewTicket()
	tic.Accept()
	tic.Prepare()
	tic.Prepare()

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

func Test_Created_To_Prepared_Is_Not_Valid(t *testing.T) {
	tic := NewTicket()

	err := tic.Prepare()
	
	if err != PrepareFromCreatedIsNotValid {
		t.Fatalf("\nshould return error")
	}
}

func Test_PrepareFromReadyToPickUpIsNotValid(t *testing.T) {
	tic := NewTicket()
	tic.Accept()
	tic.Prepare()
	tic.ReadyToPickUp()

	err := tic.Prepare()
	
	if err != PrepareFromReadyToPickUpIsNotValid {
		t.Fatalf("\nshould return error")
	}
}

//
// ReadyToPickUp
//
func Test_Prepare_to_ReadyToPickUp_Set_ReadyToPickUp_State(t *testing.T) {
	want := ReadyToPickUp
	tic := NewTicket()
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
	tic := NewTicket()
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
	tic := NewTicket()
	
	err := tic.ReadyToPickUp()

	if err != ReadyToPickUpFromCreatedIsNotValid {
		t.Fatalf("\nshould return error")
	}
}

func Test_ReadyToPickUpFromAcceptedIsNotValid(t *testing.T) {
	tic := NewTicket()
	tic.Accept()
	
	err := tic.ReadyToPickUp()

	if err != ReadyToPickUpFromAcceptedIsNotValid {
		t.Fatalf("\nshould return error")
	}
}
