package events

import (
	"errors"
	"fmt"
)

var (
	ErrDidntUseCtor error = errors.New("use constructor")
)

type Builder struct {
	valid  bool
	events events
}

func NewBuilder() Builder {
	return Builder{
		valid: true,
		events: events{
			listeners:            map[eventKey][]listener{},
			goroutinePerListener: false,
		},
	}
}

func (b Builder) GoroutinePerListener(use bool) Builder {
	if !b.valid {
		panic(fmt.Sprintf("%s\n", ErrDidntUseCtor.Error()))
	}
	b.events.goroutinePerListener = use
	return b
}

func (b Builder) Wrap(wrapped func(Builder) Builder) Builder {
	if !b.valid {
		panic(fmt.Sprintf("%s\n", ErrDidntUseCtor.Error()))
	}
	return wrapped(b)
}

func Listen[Event event](b Builder, listener func(Event)) Builder {
	if !b.valid {
		panic(fmt.Sprintf("%s\n", ErrDidntUseCtor.Error()))
	}
	eventKey := getEventKey[Event]()
	_, ok := b.events.listeners[eventKey]
	if !ok {
		b.events.listeners[eventKey] = nil
	}
	b.events.listeners[eventKey] = append(b.events.listeners[eventKey], listener)
	return b
}

func (b Builder) Build() Events {
	if !b.valid {
		panic(fmt.Sprintf("%s\n", ErrDidntUseCtor.Error()))
	}
	e := b.events
	return &e
}
