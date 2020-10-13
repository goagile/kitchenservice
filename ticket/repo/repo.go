package repo

//
// TicketRepo
//
type TicketRepo interface {
	NextID() int64
}
