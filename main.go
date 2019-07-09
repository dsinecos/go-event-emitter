package main

import (
	"fmt"
	"sync"

	"github.com/dsinecos/go-event-emitter/evtemtr"
)

type button struct {
	buttonType string
	*evtemtr.EventEmitter
}

// ButtonEvent TODO
type ButtonEvent struct {
	string
}

// GetEventName TODO
func (be ButtonEvent) GetEventName() string {
	return be.string
}

func main() {
	button := button{
		"light",
		evtemtr.New(),
	}

	onMouseOver := make(chan evtemtr.EventTuple)
	onceMouseOver := make(chan evtemtr.EventTuple)

	mouseOverButtonEvent := ButtonEvent{"mouseover"}

	var wg sync.WaitGroup

	button.On(mouseOverButtonEvent, onMouseOver).List()
	wg.Add(1)
	listen(onMouseOver, &wg)

	button.Once(mouseOverButtonEvent, onceMouseOver).List()
	// wg.Add(1)
	// listen(onceMouseOver, &wg)

	button.Remove(mouseOverButtonEvent, onceMouseOver).List()

	button.Emit(mouseOverButtonEvent, 2)

	wg.Wait()
	button.List()
}

func listen(c <-chan evtemtr.EventTuple, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		fmt.Println("Listener invoked ", <-c)
	}()
}
