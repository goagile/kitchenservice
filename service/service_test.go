package service

import (
	"os"
	"testing"

	"github.com/goagile/kitchenservice/event"
	"github.com/goagile/kitchenservice/event/bus"
	"github.com/goagile/kitchenservice/ticket"
	"github.com/goagile/kitchenservice/ticket/repo/pg"
	"github.com/goagile/kitchenservice/utils"
)

func TestMain(m *testing.M) {
	// setup()
	pg.ResetSeq("tickets_ticket_id_seq")
	pg.DeleteAll("tickets")
	DomainEvents = bus.New()
	ticket.Clock = utils.NewFakeClock(utils.TestDateTime)

	code := m.Run()

	// teardown()
	pg.DeleteAll("tickets")
	os.Exit(code)
}

//
// CreateTicket
//
func Test_CreateTicket_Publish_TicketCreated(t *testing.T) {
	want, got := true, false

	DomainEvents.AddFunc(func(e event.Event) {
		switch e.(type) {
		case event.TicketCreated:
			got = true
		}
	})

	_, err := CreateTicket(TicketDetails{OrderID: 123})
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_CreateTicket_Save_Ticket(t *testing.T) {
	want := int64(123)

	id, err := CreateTicket(TicketDetails{OrderID: want})
	if err != nil {
		t.Fatal(err)
	}

	tic, err := TicketRepo.Find(id)
	if err != nil {
		t.Fatalf("TicketRepo.Find: %v\n", err)
	}

	got := tic.OrderID

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// AcceptTicket
//
func Test_AcceptTicket_Publish_TicketAccepted(t *testing.T) {
	want, got := true, false

	DomainEvents.AddFunc(func(e event.Event) {
		switch e.(type) {
		case event.TicketAccepted:
			got = true
		}
	})

	err := AcceptTicket(ticket.TicketID(1))
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_AcceptTicket_Update_Ticket(t *testing.T) {
	want := ticket.Accepted

	id, err := CreateTicket(TicketDetails{})
	if err != nil {
		t.Fatal(err)
	}

	err = AcceptTicket(id)
	if err != nil {
		t.Fatal(err)
	}

	tic, err := TicketRepo.Find(id)
	if err != nil {
		t.Fatalf("TicketRepo.Find: %v\n", err)
	}

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// PrepareTicket
//
func Test_PrepareTicket_Publish_TicketPrepared(t *testing.T) {
	want, got := true, false

	DomainEvents.AddFunc(func(e event.Event) {
		switch e.(type) {
		case event.TicketPrepared:
			got = true
		}
	})

	err := PrepareTicket(ticket.TicketID(1))
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_PrepareTicket_Update_Ticket(t *testing.T) {
	want := ticket.Prepared

	id, err := CreateTicket(TicketDetails{})
	if err != nil {
		t.Fatal(err)
	}

	err = AcceptTicket(id)
	if err != nil {
		t.Fatal(err)
	}

	err = PrepareTicket(id)
	if err != nil {
		t.Fatal(err)
	}

	tic, err := TicketRepo.Find(id)
	if err != nil {
		t.Fatalf("TicketRepo.Find: %v\n", err)
	}

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// ReadyToPickUp
//
func Test_ReadyToPickUpTicket_Publish_TicketReadyToPickUp(t *testing.T) {
	want, got := true, false

	DomainEvents.AddFunc(func(e event.Event) {
		switch e.(type) {
		case event.TicketReadyToPickUp:
			got = true
		}
	})

	err := ReadyToPickUp(ticket.TicketID(1))
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_ReadyToPickUp_Update_Ticket(t *testing.T) {
	want := ticket.ReadyToPickUp

	id, err := CreateTicket(TicketDetails{})
	if err != nil {
		t.Fatal(err)
	}

	err = AcceptTicket(id)
	if err != nil {
		t.Fatal(err)
	}

	err = PrepareTicket(id)
	if err != nil {
		t.Fatal(err)
	}

	err = ReadyToPickUp(id)
	if err != nil {
		t.Fatal(err)
	}

	tic, err := TicketRepo.Find(id)
	if err != nil {
		t.Fatalf("TicketRepo.Find: %v\n", err)
	}

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// Cancel
//
func Test_CancelTicket_Publish_TicketCancelled(t *testing.T) {
	want, got := true, false

	DomainEvents.AddFunc(func(e event.Event) {
		switch e.(type) {
		case event.TicketCancelled:
			got = true
		}
	})

	err := Cancel(ticket.TicketID(1))
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Cancel_Update_Ticket(t *testing.T) {
	want := ticket.Cancelled

	id, err := CreateTicket(TicketDetails{})
	if err != nil {
		t.Fatal(err)
	}

	err = Cancel(id)
	if err != nil {
		t.Fatal(err)
	}

	tic, err := TicketRepo.Find(id)
	if err != nil {
		t.Fatalf("TicketRepo.Find: %v\n", err)
	}

	got := tic.State

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}
