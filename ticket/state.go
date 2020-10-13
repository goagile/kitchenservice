package ticket

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
