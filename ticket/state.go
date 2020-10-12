package ticket

//
// State
//
type State string

const (
	Created State = "Created"
	Accepted State = "Accepted"
	Prepared State = "Prepared"
	ReadyToPickUp State = "ReadyToPickUp"
	Cancelled State = "Cancelled"
)
