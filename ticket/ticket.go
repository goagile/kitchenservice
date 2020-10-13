package ticket

import (
	"time"
)

//
// TicketID
//
type TicketID int64

//
// Ticket
//
type Ticket struct {
	ID TicketID

	State State

	CreatedAt        time.Time
	CancelledAt      time.Time
	AcceptedAt       time.Time
	PreparedAt       time.Time
	ReadyForPickUpAt time.Time
}

//
// New
//
func New(id TicketID) *Ticket {
	return &Ticket{
		ID:        id,
		State:     Created,
		CreatedAt: DefaultClock.Now(),
	}
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

	tic.CancelledAt = DefaultClock.Now()

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

	tic.AcceptedAt = DefaultClock.Now()

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

	tic.PreparedAt = DefaultClock.Now()

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

	tic.ReadyForPickUpAt = DefaultClock.Now()

	return nil
}
