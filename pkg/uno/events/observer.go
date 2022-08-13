package events

import "github.com/glugate/uno/pkg/uno/log"

// Observer is a behavioral design pattern that lets you define
// a subscription mechanism to notify multiple objects about any
// events that happen to the object they’re observing.
type Observer struct {
	listeners []EventHandler
	logger    *log.Log
}

// NewObserver creates new Observer object
func NewObserver() *Observer {
	return &Observer{
		logger: log.DefaultLogFactory().NewLogger(),
	}
}

// AddListener is usually called from objects that would like to listen for some
// events and therefore have some handler function
func (o *Observer) AddListener(eventName string, h EventHandler) {
	o.listeners = append(o.listeners, h)
}

// Dispatch emits particular event and so all listeners are notified.
func (o *Observer) Dispatch(eventName string) {
	e := NewEvent(eventName)
	o.logger.Event(eventName)
	for _, h := range o.listeners {
		h(e)
	}
}
