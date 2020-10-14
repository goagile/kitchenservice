package pg

import (
	"os"
	"testing"

	"github.com/goagile/kitchenservice/ticket"
	"github.com/goagile/kitchenservice/utils"
)

func init() {
	DB = Connected(os.Getenv("KITCHEN_PG"))
	ticket.Clock = utils.NewFakeClock(utils.TestDateTime)
}

func TestMain(m *testing.M) {
	ResetSeq("tickets_ticket_id_seq")
	DeleteAll("tickets")
	os.Exit(m.Run())
}

//
// NextID
//
func Test_NextID(t *testing.T) {
	want := ticket.TicketID(1)

	got, err := TicketRepo.NextID()
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
	want := true

	id, err := TicketRepo.NextID()
	if err != nil {
		t.Fatal(err)
	}

	a := ticket.New(id)
	a.Accept()
	a.Prepare()
	a.ReadyToPickUp()
	a.Cancel()

	err = TicketRepo.Save(a)
	if err != nil {
		t.Fatal(err)
	}

	b, err := TicketRepo.Find(id)
	if err != nil {
		t.Fatal(err)
	}

	got := a.Eq(b)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Save_As_Update(t *testing.T) {
	want := true

	id, err := TicketRepo.NextID()
	if err != nil {
		t.Fatal(err)
	}

	a := ticket.New(id)

	err = TicketRepo.Save(a)
	if err != nil {
		t.Fatal(err)
	}

	b, err := TicketRepo.Find(id)
	if err != nil {
		t.Fatal(err)
	}
	b.Accept()
	b.Prepare()
	b.ReadyToPickUp()
	b.Cancel()

	err = TicketRepo.Save(b)
	if err != nil {
		t.Fatal(err)
	}

	c, err := TicketRepo.Find(id)
	if err != nil {
		t.Fatal(err)
	}

	got := c.Eq(b)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}
