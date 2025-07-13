package events

import (
	"reflect"
)

type eventKey reflect.Type

type event interface{}

type listener func(emiter Events, event any)
type anyListener func(emiter Events, event any)

type events struct {
	listeners            map[eventKey][]listener
	allListeners         []anyListener
	goroutinePerListener bool
}

type Events *events

func getEventKey[Event event]() eventKey {
	return reflect.TypeFor[Event]()
}

func getAnyEventKey(event any) eventKey {
	return reflect.TypeOf(event)
}

func emitAny(e Events, event any) {
	if e.goroutinePerListener {
		for _, listener := range e.allListeners {
			go listener(e, event)
		}
	} else {
		for _, listener := range e.allListeners {
			listener(e, event)
		}
	}
}

func Emit[Event event](e Events, event Event) {
	eventKey := getEventKey[Event]()
	eventListeners, ok := e.listeners[eventKey]
	emitAny(e, event)
	if !ok {
		return
	}
	if e.goroutinePerListener {
		for _, listener := range eventListeners {
			go listener(e, event)
		}
	} else {
		for _, listener := range eventListeners {
			listener(e, event)
		}
	}
}

func EmitAny(e Events, event any) {
	eventKey := getAnyEventKey(event)
	eventListeners, ok := e.listeners[eventKey]
	emitAny(e, event)
	if !ok {
		return
	}
	if e.goroutinePerListener {
		for _, listener := range eventListeners {
			go listener(e, event)
		}
	} else {
		for _, listener := range eventListeners {
			listener(e, event)
		}
	}
}
