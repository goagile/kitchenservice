package pg

import (
	"database/sql"

	"github.com/goagile/kitchenservice/ticket"
	_ "github.com/lib/pq"
)

var DB *sql.DB

//
// ticketRepo
//
type ticketRepo struct{}

//
// NewTicketRepo
//
func NewTicketRepo() *ticketRepo {
	return &ticketRepo{}
}

//
// NextID
//
func (r *ticketRepo) NextID() (ticket.TicketID, error) {
	var id ticket.TicketID
	row := DB.QueryRow("SELECT nextval('tickets_id_seq')")
	if err := row.Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

//
// Save
//
const insertQuery = `
	INSERT INTO tickets(ticket_id, state)
	VALUES($1, $2)
`

func (r *ticketRepo) Save(tic *ticket.Ticket) error {
	_, err := DB.Exec(insertQuery, tic.ID, tic.State)
	return err
}
