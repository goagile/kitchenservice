package service

import (
	"os"

	"github.com/goagile/kitchenservice/event"
	"github.com/goagile/kitchenservice/event/bus"
	"github.com/goagile/kitchenservice/ticket"
	"github.com/goagile/kitchenservice/ticket/repo"
	"github.com/goagile/kitchenservice/ticket/repo/pg"
)

//
// DomainEvents
//
var DomainEvents bus.Bus

//
// TicketRepo
//
var TicketRepo repo.TicketRepo

func init() {
	pg.DB = pg.Connected(os.Getenv("KITCHEN_PG"))
	TicketRepo = pg.NewTicketRepo()
	DomainEvents = bus.New()
}

//
// TicketDetails
//
type TicketDetails struct {
	OrderID int64 `json:"order_id"`
}

//
// CreateTicket
//
func CreateTicket(details TicketDetails) (ticket.TicketID, error) {
	id, err := TicketRepo.NextID()
	if err != nil {
		return id, err
	}

	tic := ticket.New(id)
	tic.OrderID = details.OrderID
	if err := TicketRepo.Save(tic); err != nil {
		return id, err
	}

	DomainEvents.Publish(event.TicketCreated{
		TicketID: tic.ID,
		OrderID:  tic.OrderID,
	})

	return id, nil
}

//
// AcceptTicket
//
func AcceptTicket(id ticket.TicketID) error {
	tic, err := TicketRepo.Find(id)
	if err != nil {
		return err
	}

	err = tic.Accept()
	if err != nil {
		return err
	}

	err = TicketRepo.Save(tic)
	if err != nil {
		return err
	}

	DomainEvents.Publish(event.TicketAccepted{
		TicketID: tic.ID,
	})

	return nil
}

//
// PrepareTicket
//
func PrepareTicket(id ticket.TicketID) error {
	tic, err := TicketRepo.Find(id)
	if err != nil {
		return err
	}

	err = tic.Prepare()
	if err != nil {
		return err
	}

	err = TicketRepo.Save(tic)
	if err != nil {
		return err
	}

	DomainEvents.Publish(event.TicketPrepared{
		TicketID: tic.ID,
	})

	return nil
}

//
// ReadyToPickUp
//
func ReadyToPickUp(id ticket.TicketID) error {
	tic, err := TicketRepo.Find(id)
	if err != nil {
		return err
	}

	err = tic.ReadyToPickUp()
	if err != nil {
		return err
	}

	err = TicketRepo.Save(tic)
	if err != nil {
		return err
	}

	DomainEvents.Publish(event.TicketReadyToPickUp{
		TicketID: tic.ID,
	})

	return nil
}

//
// Cancel
//
func Cancel(id ticket.TicketID) error {
	tic, err := TicketRepo.Find(id)
	if err != nil {
		return err
	}

	err = tic.Cancel()
	if err != nil {
		return err
	}

	err = TicketRepo.Save(tic)
	if err != nil {
		return err
	}

	DomainEvents.Publish(event.TicketCancelled{
		TicketID: tic.ID,
	})

	return nil
}
