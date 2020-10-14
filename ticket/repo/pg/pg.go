package pg

import (
	"database/sql"

	"github.com/goagile/kitchenservice/ticket"
	"github.com/goagile/kitchenservice/ticket/repo"
	_ "github.com/lib/pq"
)

var DB *sql.DB

var TicketRepo repo.TicketRepo

func init() {
	TicketRepo = NewTicketRepo()
}

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
const nextIDQuery = `
	SELECT nextval('tickets_ticket_id_seq')
`

func (r *ticketRepo) NextID() (ticket.TicketID, error) {
	var id ticket.TicketID
	row := DB.QueryRow(nextIDQuery)
	if err := row.Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

//
// Save
//
const saveQuery = `
	INSERT INTO tickets(
		ticket_id,
		state, 
		created_at,
		accepted_at,
		prepared_at,
		ready_to_pickup_at,
		cancelled_at,
		order_id
	)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8)

	ON CONFLICT(ticket_id) 
	
	DO UPDATE SET 
		ticket_id=$1,
		state=$2, 
		created_at=$3,
		accepted_at=$4,
		prepared_at=$5,
		ready_to_pickup_at=$6,
		cancelled_at=$7,
		order_id=$8
`

func (r *ticketRepo) Save(tic *ticket.Ticket) error {
	_, err := DB.Exec(saveQuery,
		tic.ID,
		tic.State,
		tic.CreatedAt,
		tic.AcceptedAt,
		tic.PreparedAt,
		tic.ReadyToPickUpAt,
		tic.CancelledAt,
		tic.OrderID,
	)
	return err
}

//
// Find
//
const findQuery = `
	SELECT 
		state, 
		created_at,
		accepted_at,
		prepared_at,
		ready_to_pickup_at,
		cancelled_at,
		order_id
	FROM tickets 
	WHERE ticket_id = $1
`

func (r *ticketRepo) Find(id ticket.TicketID) (*ticket.Ticket, error) {
	tic := ticket.New(id)
	row := DB.QueryRow(findQuery, id)
	err := row.Scan(
		&tic.State,
		&tic.CreatedAt,
		&tic.AcceptedAt,
		&tic.PreparedAt,
		&tic.ReadyToPickUpAt,
		&tic.CancelledAt,
		&tic.OrderID,
	)
	if err != nil {
		return nil, err
	}
	return tic, nil
}
