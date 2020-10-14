package bus

import "github.com/goagile/kitchenservice/event"

//
// Bus
//
type Bus struct {
	handlers []Handler
}

//
// Publish
//
func (b *Bus) Publish(e event.Event) {
	for _, h := range b.handlers {
		h.Handle(e)
	}
}

//
// Add
//
func (b *Bus) Add(h Handler) {
	b.handlers = append(b.handlers, h)
}

//
// AddFunc
//
func (b *Bus) AddFunc(f func(e event.Event)) {
	b.Add(&funcHandler{f})
}

//
// Handler
//
type Handler interface {
	Handle(e event.Event)
}

//
// funcHandler
//
type funcHandler struct {
	f func(e event.Event)
}

//
// Handle
//
func (h *funcHandler) Handle(e event.Event) {
	h.f(e)
}
