package bus

import "github.com/goagile/kitchenservice/event"

//
// bus
//
type Bus interface {
	Publish(e event.Event)
	Add(h Handler)
	AddFunc(f func(e event.Event))
}

type bus struct {
	handlers []Handler
}

//
// New
//
func New() *bus {
	return &bus{
		handlers: make([]Handler, 0),
	}
}

//
// Publish
//
func (b *bus) Publish(e event.Event) {
	for _, h := range b.handlers {
		h.Handle(e)
	}
}

//
// Add
//
func (b *bus) Add(h Handler) {
	b.handlers = append(b.handlers, h)
}

//
// AddFunc
//
func (b *bus) AddFunc(f func(e event.Event)) {
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
