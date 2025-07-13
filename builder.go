package events

type builder struct{}

type Builder struct {
	events *events
}

func NewBuilder() Builder {
	return Builder{
		events: &events{
			listeners:            map[eventKey][]listener{},
			allListeners:         []anyListener{},
			goroutinePerListener: false,
		},
	}
}

func (b Builder) GoroutinePerListener(use bool) {
	b.events.goroutinePerListener = use
}

func Listen[Event event](b Builder, listener func(emiter Events, event Event)) {
	eventKey := getEventKey[Event]()
	_, ok := b.events.listeners[eventKey]
	if !ok {
		b.events.listeners[eventKey] = nil
	}
	b.events.listeners[eventKey] = append(b.events.listeners[eventKey], func(emiter Events, e any) {
		listener(emiter, e.(Event))
	})
}

func ListenToAll(b Builder, listener func(emiter Events, event any)) {
	b.events.allListeners = append(b.events.allListeners, listener)
}

func (b Builder) Build() Events {
	e := *b.events
	return &e
}
