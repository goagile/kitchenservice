package service

import (
	"os"
	"testing"

	"github.com/goagile/kitchenservice/event"
	"github.com/goagile/kitchenservice/ticket"
	"github.com/goagile/kitchenservice/ticket/repo/pg"
	"github.com/goagile/kitchenservice/utils"
)

func init() {
	pg.DB = pg.Connected(os.Getenv("KITCHEN_PG"))
}

func TestMain(m *testing.M) {
	// setup()
	pg.ResetSeq("tickets_ticket_id_seq")
	pg.DeleteAll("tickets")
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

	DomainEvents := &fakePublisher{}
	TicketRepo := pg.NewTicketRepo()
	Kitchen := NewService(DomainEvents, "test", TicketRepo)

	_, err := Kitchen.CreateTicket(TicketDetails{OrderID: 123})
	if err != nil {
		t.Fatal(err)
	}

	got = DomainEvents.Raised

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_CreateTicket_Save_Ticket(t *testing.T) {
	want := int64(123)

	DomainEvents := &fakePublisher{}
	TicketRepo := pg.NewTicketRepo()
	Kitchen := NewService(DomainEvents, "test", TicketRepo)

	id, err := Kitchen.CreateTicket(TicketDetails{OrderID: want})
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

	DomainEvents := &fakePublisher{}
	TicketRepo := pg.NewTicketRepo()
	Kitchen := NewService(DomainEvents, "test", TicketRepo)

	err := Kitchen.AcceptTicket(ticket.TicketID(1))
	if err != nil {
		t.Fatal(err)
	}

	got = DomainEvents.Raised

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_AcceptTicket_Update_Ticket(t *testing.T) {
	want := ticket.Accepted

	DomainEvents := &fakePublisher{}
	TicketRepo := pg.NewTicketRepo()
	Kitchen := NewService(DomainEvents, "test", TicketRepo)

	id, err := Kitchen.CreateTicket(TicketDetails{})
	if err != nil {
		t.Fatal(err)
	}

	err = Kitchen.AcceptTicket(id)
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

	DomainEvents := &fakePublisher{}
	TicketRepo := pg.NewTicketRepo()
	Kitchen := NewService(DomainEvents, "test", TicketRepo)

	err := Kitchen.PrepareTicket(ticket.TicketID(1))
	if err != nil {
		t.Fatal(err)
	}

	got = DomainEvents.Raised

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_PrepareTicket_Update_Ticket(t *testing.T) {
	want := ticket.Prepared

	DomainEvents := &fakePublisher{}
	TicketRepo := pg.NewTicketRepo()
	Kitchen := NewService(DomainEvents, "test", TicketRepo)

	id, err := Kitchen.CreateTicket(TicketDetails{})
	if err != nil {
		t.Fatal(err)
	}

	err = Kitchen.AcceptTicket(id)
	if err != nil {
		t.Fatal(err)
	}

	err = Kitchen.PrepareTicket(id)
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

	DomainEvents := &fakePublisher{}
	TicketRepo := pg.NewTicketRepo()
	Kitchen := NewService(DomainEvents, "test", TicketRepo)

	err := Kitchen.ReadyToPickUp(ticket.TicketID(1))
	if err != nil {
		t.Fatal(err)
	}

	got = DomainEvents.Raised

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_ReadyToPickUp_Update_Ticket(t *testing.T) {
	want := ticket.ReadyToPickUp

	DomainEvents := &fakePublisher{}
	TicketRepo := pg.NewTicketRepo()
	Kitchen := NewService(DomainEvents, "test", TicketRepo)

	id, err := Kitchen.CreateTicket(TicketDetails{})
	if err != nil {
		t.Fatal(err)
	}

	err = Kitchen.AcceptTicket(id)
	if err != nil {
		t.Fatal(err)
	}

	err = Kitchen.PrepareTicket(id)
	if err != nil {
		t.Fatal(err)
	}

	err = Kitchen.ReadyToPickUp(id)
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

	DomainEvents := &fakePublisher{}
	TicketRepo := pg.NewTicketRepo()
	Kitchen := NewService(DomainEvents, "test", TicketRepo)

	err := Kitchen.Cancel(ticket.TicketID(1))
	if err != nil {
		t.Fatal(err)
	}

	got = DomainEvents.Raised

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Cancel_Update_Ticket(t *testing.T) {
	want := ticket.Cancelled

	DomainEvents := &fakePublisher{}
	TicketRepo := pg.NewTicketRepo()
	Kitchen := NewService(DomainEvents, "test", TicketRepo)

	id, err := Kitchen.CreateTicket(TicketDetails{})
	if err != nil {
		t.Fatal(err)
	}

	err = Kitchen.Cancel(id)
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

type fakePublisher struct {
	Raised bool
}

func (p *fakePublisher) Publish(topic string, e event.Event) {
	p.Raised = true
}
