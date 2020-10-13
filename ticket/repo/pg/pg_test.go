package pg

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

func init() {
	connect()
	resetTicketsIDs()
}

func connect() {
	var err error
	DB, err = sql.Open("postgres", os.Getenv("KITCHEN_PG"))
	if err != nil {
		log.Fatalf("DB Open err: %v", err)
	}
}

func resetTicketsIDs() {
	q := "ALTER SEQUENCE tickets_id_seq RESTART WITH 1"
	if _, err := DB.Exec(q); err != nil {
		log.Fatalf("DB Reset Sequence err: %v", err)
	}
}

//
// NextID
//
func Test_NextID(t *testing.T) {
	want := int64(1)
	r := NewTicketRepo()

	got, err := r.NextID()
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}
