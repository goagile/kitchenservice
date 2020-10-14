package service

import (
	"os"
	"testing"

	"github.com/goagile/kitchenservice/event"
	"github.com/goagile/kitchenservice/event/bus"
	"github.com/goagile/kitchenservice/ticket"
	"github.com/goagile/kitchenservice/ticket/repo/pg"
)

func TestMain(m *testing.M) {
	pg.ResetSeq("tickets_ticket_id_seq")
	pg.DeleteAll("tickets")
	DomainEvents = bus.New()
	os.Exit(m.Run())
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

	err := CreateTicket(TicketDetails{OrderID: 123})
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_CreateTicket_Save_Ticket(t *testing.T) {
	want := int64(123)

	var id ticket.TicketID

	DomainEvents.AddFunc(func(e event.Event) {
		switch v := e.(type) {
		case event.TicketCreated:
			id = v.TicketID
		}
	})

	err := CreateTicket(TicketDetails{OrderID: int64(123)})
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