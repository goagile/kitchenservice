package event

import (
	"github.com/goagile/kitchenservice/ticket"
)

//
// Event
//
type Event interface{}

//
// TicketCreated
//
type TicketCreated struct {
	TicketID ticket.TicketID
	OrderID  int64
}

//
// TicketAccepted
//
type TicketAccepted struct {
	TicketID ticket.TicketID
}

//
// TicketPrepared
//
type TicketPrepared struct {
	TicketID ticket.TicketID
}

//
// TicketReadyToPickUp
//
type TicketReadyToPickUp struct {
	TicketID ticket.TicketID
}

//
// TicketCancelled
//
type TicketCancelled struct {
	TicketID ticket.TicketID
}
