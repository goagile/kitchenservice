package pg

import (
	"database/sql"

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
func (r *ticketRepo) NextID() (int64, error) {
	var id int64
	row := DB.QueryRow("SELECT nextval('tickets_id_seq')")
	if err := row.Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}
