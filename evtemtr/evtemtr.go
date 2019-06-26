package evtemtr

import (
	"fmt"
)

// EventEmitter TODO
type EventEmitter struct {
	eventQueue map[string][]chan<- string
}

// New TODO
func New() *EventEmitter {
	newEventEmitter := &EventEmitter{
		eventQueue: make(map[string][]chan<- string),
	}
	return newEventEmitter
}

// On TODO
func (emtr *EventEmitter) On(event string, listener chan<- string) *EventEmitter {
	if _, ok := emtr.eventQueue[event]; !ok {
		emtr.eventQueue[event] = make([]chan<- string, 0)
	}
	emtr.eventQueue[event] = append(emtr.eventQueue[event], listener)

	return emtr
}

// Emit TODO
func (emtr *EventEmitter) Emit(event string) *EventEmitter {

	if _, ok := emtr.eventQueue[event]; !ok {
		fmt.Printf("No listeners attached to the event %s\n", event)
		return emtr
	}

	for _, listener := range emtr.eventQueue[event] {
		go func(listener chan<- string) {
			listener <- event
		}(listener)
	}

	return emtr
}

// Remove TODO
func (emtr *EventEmitter) Remove(event string, listener chan<- string) *EventEmitter {

	if _, ok := emtr.eventQueue[event]; !ok {
		fmt.Printf("No listeners attached to the event %s\n", event)
		return emtr
	}

	if queuedListeners, ok := emtr.eventQueue[event]; ok {
		for index, queuedlistener := range queuedListeners {
			if listener == queuedlistener {
				emtr.eventQueue[event] = append(emtr.eventQueue[event][:index], emtr.eventQueue[event][index+1:]...)
				break
			}
		}
	}
	return emtr
}

// List TODO
func (emtr *EventEmitter) List() {
	for event, listeners := range emtr.eventQueue {
		fmt.Printf("%v, %v listeners \n", event, len(listeners))
	}
}
