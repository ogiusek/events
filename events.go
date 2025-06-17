package events

import (
	"reflect"
)

type eventKey reflect.Type

type event interface{}

type listener any

type events struct {
	listeners            map[eventKey][]listener
	goroutinePerListener bool
}

type Events *events

func getEventKey[Event event]() eventKey {
	return reflect.TypeFor[Event]()
}

func Emit[Event event](e Events, event Event) {
	eventKey := getEventKey[Event]()
	eventListeners, ok := e.listeners[eventKey]
	if !ok {
		return
	}
	if e.goroutinePerListener {
		for _, listener := range eventListeners {
			go listener.(func(Event))(event)
		}
	} else {
		for _, listener := range eventListeners {
			listener.(func(Event))(event)
		}
	}
}
