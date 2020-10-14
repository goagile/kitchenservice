package ticket

import (
	"errors"
	"fmt"
	"time"

	"github.com/goagile/kitchenservice/utils"
)

//
// Ticket
//
type Ticket struct {
	ID              TicketID
	OrderID         int64
	State           State
	CreatedAt       time.Time
	CancelledAt     time.Time
	AcceptedAt      time.Time
	PreparedAt      time.Time
	ReadyToPickUpAt time.Time
}

//
// TicketID
//
type TicketID int64

//
// State
//
type State string

const (
	Created       State = "CREATED"
	Accepted      State = "ACCEPTED"
	Prepared      State = "PREPARED"
	ReadyToPickUp State = "READY_FOR_PICKUP"
	Cancelled     State = "CANCELLED"
)

//
// New
//
func New(id TicketID) *Ticket {
	return &Ticket{
		ID:        id,
		State:     Created,
		CreatedAt: Clock.Now(),
	}
}

//
// String
//
func (tic *Ticket) String() string {
	return fmt.Sprintf(
		"Ticket:\n"+
			"\tID:%v\n"+
			"\tOrderID:%v\n"+
			"\tState:%v\n"+
			"\tCreatedAt:%v\n"+
			"\tPreparedAt:%v\n"+
			"\tReadyToPickUpAt:%v\n"+
			"\tCancelledAt:%v\n",
		tic.ID,
		tic.OrderID,
		tic.State,
		tic.CreatedAt,
		tic.PreparedAt,
		tic.ReadyToPickUpAt,
		tic.CancelledAt,
	)
}

//
// Eq
//
func (tic *Ticket) Eq(other *Ticket) bool {
	return tic.ID == other.ID &&
		tic.State == other.State &&
		utils.DateTimeEq(tic.AcceptedAt, other.AcceptedAt) &&
		utils.DateTimeEq(tic.PreparedAt, other.PreparedAt) &&
		utils.DateTimeEq(tic.ReadyToPickUpAt, other.ReadyToPickUpAt) &&
		utils.DateTimeEq(tic.CancelledAt, other.CancelledAt) &&
		tic.OrderID == other.OrderID
}

//
// Cancel
//
func (tic *Ticket) Cancel() error {
	switch tic.State {

	case Created, ReadyToPickUp:
		tic.State = Cancelled

	case Accepted:
		return CancelFromAcceptedIsNotValid

	case Prepared:
		return CancelFromPreparedIsNotValid
	}

	tic.CancelledAt = Clock.Now()

	return nil
}

//
// Accept
//
func (tic *Ticket) Accept() error {
	switch tic.State {

	case Created, Accepted:
		tic.State = Accepted

	case Prepared:
		return AcceptedFromPreparedIsNotValid

	case ReadyToPickUp:
		return AcceptedFromReadyToPickUpIsNotValid

	}

	tic.AcceptedAt = Clock.Now()

	return nil
}

//
// Prepare
//
func (tic *Ticket) Prepare() error {
	switch tic.State {

	case Created:
		return PrepareFromCreatedIsNotValid

	case Accepted, Prepared:
		tic.State = Prepared

	case ReadyToPickUp:
		return PrepareFromReadyToPickUpIsNotValid
	}

	tic.PreparedAt = Clock.Now()

	return nil
}

//
// ReadyToPickUp
//
func (tic *Ticket) ReadyToPickUp() error {
	switch tic.State {

	case Created:
		return ReadyToPickUpFromCreatedIsNotValid

	case Accepted:
		return ReadyToPickUpFromAcceptedIsNotValid

	case Prepared, ReadyToPickUp:
		tic.State = ReadyToPickUp
	}

	tic.ReadyToPickUpAt = Clock.Now()

	return nil
}

//
// Errrors
//
var (
	AcceptedFromPreparedIsNotValid      = errors.New("AcceptedFromPreparedIsNotValid")
	AcceptedFromReadyToPickUpIsNotValid = errors.New("AcceptedFromReadyToPickUpIsNotValid")

	PrepareFromCreatedIsNotValid       = errors.New("PrepareFromCreatedIsNotValid")
	PrepareFromReadyToPickUpIsNotValid = errors.New("PrepareFromReadyToPickUpIsNotValid")

	ReadyToPickUpFromCreatedIsNotValid  = errors.New("ReadyToPickUpFromCreatedIsNotValid")
	ReadyToPickUpFromAcceptedIsNotValid = errors.New("ReadyToPickUpFromAcceptedIsNotValid")

	CancelFromAcceptedIsNotValid = errors.New("CancelFromAcceptedIsNotValid")
	CancelFromPreparedIsNotValid = errors.New("CancelFromPreparedIsNotValid")
)

//
// Clock
//
var Clock utils.Clock = utils.NewSystemClock()
