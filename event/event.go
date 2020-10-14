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
}
