package repo

import "github.com/goagile/kitchenservice/ticket"

//
// TicketRepo
//
type TicketRepo interface {
	NextID() ticket.TicketID
	Save() error
	Find(ticket.TicketID) (*ticket.Ticket, error)
}
