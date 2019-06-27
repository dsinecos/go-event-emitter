package evtemtr

// Event TODO
type Event interface {
	GetEventName() string
}

// EventData TODO
type EventData interface{}

// EventTuple TODO
type EventTuple struct {
	Event
	EventData
}
