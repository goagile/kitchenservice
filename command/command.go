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
// AcceptTicket
//
type AcceptTicket struct {
	TicketID ticket.TicketID `json:"ticket_id"`
}

//
// CancelTicket
//
type CancelTicket struct {
	TicketID ticket.TicketID `json:"ticket_id"`
}
