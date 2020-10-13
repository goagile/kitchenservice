package pg

import (
	"os"
	"testing"

	"github.com/goagile/kitchenservice/ticket"
	"github.com/goagile/kitchenservice/utils"
)

var repo = NewTicketRepo()

func init() {
	DB = Connected(os.Getenv("KITCHEN_PG"))
	ResetSeq("tickets_id_seq")
	DeleteAll("tickets")
	ticket.DefaultClock = utils.NewFakeClock(utils.TestDateTime)
}

//
// NextID
//
func Test_NextID(t *testing.T) {
	ResetSeq("tickets_id_seq")
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
	ResetSeq("tickets_id_seq")

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
	ResetSeq("tickets_id_seq")

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
