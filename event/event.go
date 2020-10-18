package event

import (
	"github.com/goagile/kitchenservice/ticket"
)

//
// Event
//
type Event interface {
	Name() string
}

//
// TicketCreated
//
type TicketCreated struct {
	TicketID ticket.TicketID `json:"ticket_id"`
	OrderID  int64           `json:"order_id"`
}

//
// Name
//
func (e *TicketCreated) Name() string {
	return "ticket_created"
}

//
// TicketAccepted
//
type TicketAccepted struct {
	TicketID ticket.TicketID `json:"ticket_id"`
}

//
// Name
//
func (e *TicketAccepted) Name() string {
	return "ticket_accepted"
}

//
// TicketPrepared
//
type TicketPrepared struct {
	TicketID ticket.TicketID `json:"ticket_id"`
}

//
// Name
//
func (e *TicketPrepared) Name() string {
	return "ticket_prepared"
}

//
// TicketReadyToPickUp
//
type TicketReadyToPickUp struct {
	TicketID ticket.TicketID `json:"ticket_id"`
}

//
// Name
//
func (e *TicketReadyToPickUp) Name() string {
	return "ticket_ready_to_pickup"
}

//
// TicketCancelled
//
type TicketCancelled struct {
	TicketID ticket.TicketID `json:"ticket_id"`
}

//
// Name
//
func (e *TicketCancelled) Name() string {
	return "ticket_cancelled"
}
