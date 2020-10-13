package pg

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/goagile/kitchenservice/ticket"
)

var repo = NewTicketRepo()

func init() {
	connect()
	resetTicketsIDs()
	removeAllTickets()
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
	// want := TicketID(1)

	id, _ := repo.NextID()
	tic := ticket.New(id)

	err := repo.Save(tic)
	if err != nil {
		t.Fatal(err)
	}

	// if want != got {
	// 	t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	// }
}
