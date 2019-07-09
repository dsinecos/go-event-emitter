package evtemtr

import (
	"fmt"
)

// EventEmitter TODO
type EventEmitter struct {
	eventQueue     map[Event][]chan<- EventTuple
	eventOnceQueue map[Event][]chan<- EventTuple
}

// New TODO
func New() *EventEmitter {
	newEventEmitter := &EventEmitter{
		eventQueue:     make(map[Event][]chan<- EventTuple),
		eventOnceQueue: make(map[Event][]chan<- EventTuple),
	}
	return newEventEmitter
}

// On TODO
func (emtr *EventEmitter) On(event Event, listener chan<- EventTuple) *EventEmitter {
	if _, ok := emtr.eventQueue[event]; !ok {
		emtr.eventQueue[event] = make([]chan<- EventTuple, 0)
	}
	emtr.eventQueue[event] = append(emtr.eventQueue[event], listener)

	return emtr
}

// Once TODO
func (emtr *EventEmitter) Once(event Event, listener chan<- EventTuple) *EventEmitter {
	if _, ok := emtr.eventOnceQueue[event]; !ok {
		emtr.eventOnceQueue[event] = make([]chan<- EventTuple, 0)
	}
	emtr.eventOnceQueue[event] = append(emtr.eventOnceQueue[event], listener)

	return emtr
}

// Emit TODO
func (emtr *EventEmitter) Emit(event Event, eventData EventData) *EventEmitter {

	_, isInEventQueue := emtr.eventQueue[event]
	_, isInEventOnceQueue := emtr.eventOnceQueue[event]

	if !isInEventQueue && !isInEventOnceQueue {
		fmt.Printf("No listeners attached to the event %s\n", event)
		return emtr
	}

	if isInEventQueue {
		for _, listener := range emtr.eventQueue[event] {
			go func(listener chan<- EventTuple) {
				listener <- EventTuple{event, eventData}
			}(listener)
		}
	}

	if isInEventOnceQueue {
		for _, listener := range emtr.eventOnceQueue[event] {
			go func(listener chan<- EventTuple) {
				listener <- EventTuple{event, eventData}
			}(listener)
		}
		delete(emtr.eventOnceQueue, event)
	}

	return emtr
}

// Remove TODO
func (emtr *EventEmitter) Remove(event Event, listener chan<- EventTuple) *EventEmitter {

	_, isInEventQueue := emtr.eventQueue[event]
	_, isInEventOnceQueue := emtr.eventOnceQueue[event]

	if !isInEventQueue && !isInEventOnceQueue {
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

	if queuedListeners, ok := emtr.eventOnceQueue[event]; ok {
		for index, queuedlistener := range queuedListeners {
			if listener == queuedlistener {
				emtr.eventOnceQueue[event] = append(emtr.eventOnceQueue[event][:index], emtr.eventOnceQueue[event][index+1:]...)
				break
			}
		}
	}

	return emtr
}

// List TODO
func (emtr *EventEmitter) List() {
	for event, listeners := range emtr.eventQueue {
		fmt.Printf("%v, %v On listeners \n", event, len(listeners))
	}

	for event, listeners := range emtr.eventOnceQueue {
		fmt.Printf("%v, %v Once listeners \n", event, len(listeners))
	}
}
