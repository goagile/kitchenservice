package event

type Event interface{}

type EmptyEvent struct{}

type Bus struct {
	handlers []Handler
}

func (b *Bus) Publish(e Event) {
	for _, h := range b.handlers {
		h.Handle(e)
	}
}

func (b *Bus) AddHandler(h Handler) {
	b.handlers = append(b.handlers, h)
}

func (b *Bus) AddFuncHandler(f func(e Event)) {
	b.AddHandler(&funcHandler{f})
}

type Handler interface {
	Handle(e Event)
}

type funcHandler struct {
	f func(e Event)
}

func (h *funcHandler) Handle(e Event) {
	h.f(e)
}
