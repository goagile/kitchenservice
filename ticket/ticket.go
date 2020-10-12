package ticket

//
// Ticket
//
type Ticket struct {
	State State
}

func NewTicket() *Ticket {
	return &Ticket{State: Created}
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

	return nil
}

func (tic *Ticket) ReadyToPickUp() error {
	switch tic.State {
	
	case Created:
		return ReadyToPickUpFromCreatedIsNotValid
	
	case Accepted:
		return ReadyToPickUpFromAcceptedIsNotValid

	case Prepared, ReadyToPickUp:
		tic.State = ReadyToPickUp
	}

	return nil
}
