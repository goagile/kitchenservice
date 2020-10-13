package pg

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/goagile/kitchenservice/ticket"
	"github.com/goagile/kitchenservice/utils"
)

var repo = NewTicketRepo()

var TestDateTime = time.Date(2020, time.October, 13, 23, 30, 10, 0, time.UTC)

func init() {
	connect()
	resetTicketsIDs()
	removeAllTickets()
	ticket.DefaultClock = utils.NewFakeClock(TestDateTime)
}

func connect() {
	var err error
	DB, err = sql.Open("postgres", os.Getenv("KITCHEN_PG"))
	if err != nil {
		log.Fatalf("DB Open err: %v", err)
	}
}

const alterSeq = "ALTER SEQUENCE tickets_id_seq RESTART WITH 1"

func resetTicketsIDs() {
	if _, err := DB.Exec(alterSeq); err != nil {
		log.Fatalf("DB Reset Sequence err: %v", err)
	}
}

const deleteAllTickets = "DELETE FROM tickets"

func removeAllTickets() {
	if _, err := DB.Exec(deleteAllTickets); err != nil {
		log.Fatalf("DB Remove All Tickets err: %v", err)
	}
}

//
// NextID
//
func Test_NextID(t *testing.T) {
	resetTicketsIDs()
	want := ticket.TicketID(1)

	got, err := repo.NextID()
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// Save
//
func Test_Save_As_Insert(t *testing.T) {
	resetTicketsIDs()

	want := true

	id, _ := repo.NextID()
	a := ticket.New(id)
	a.Accept()
	a.Prepare()
	a.ReadyToPickUp()
	a.Cancel()

	if err := repo.Save(a); err != nil {
		t.Fatal(err)
	}

	b, err := repo.Find(id)
	if err != nil {
		t.Fatal(err)
	}

	got := a.Eq(b)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Save_As_Update(t *testing.T) {
	resetTicketsIDs()

	want := true

	id, _ := repo.NextID()
	a := ticket.New(id)

	if err := repo.Save(a); err != nil {
		t.Fatal(err)
	}

	b, err := repo.Find(id)
	if err != nil {
		t.Fatal(err)
	}
	b.Accept()
	b.Prepare()
	b.ReadyToPickUp()
	b.Cancel()

	if err := repo.Save(b); err != nil {
		t.Fatal(err)
	}

	c, err := repo.Find(id)
	if err != nil {
		t.Fatal(err)
	}

	got := c.Eq(b)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}
