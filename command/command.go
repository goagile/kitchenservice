package command

import (
	"github.com/goagile/kitchenservice/ticket"
)

//
// CreateTicket
//
type CreateTicket struct {
	OrderID int64 `json:"order_id"`
}

//
// Name
//
func (c *CreateTicket) Name() string {
	return "create_ticket"
}

//
// AcceptTicket
//
type AcceptTicket struct {
	TicketID ticket.TicketID `json:"ticket_id"`
}

//
// Name
//
func (c *AcceptTicket) Name() string {
	return "accept_ticket"
}
