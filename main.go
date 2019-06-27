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

	onClick1 := make(chan string)
	onClick2 := make(chan string)

	clickButtonEvent := ButtonEvent{"click"}
	mouseOverButtonEvent := ButtonEvent{"mouseover"}

	var wg sync.WaitGroup
	button.On(clickButtonEvent, onClick1).List()
	wg.Add(1)
	listen(onClick1, &wg)
	button.On(clickButtonEvent, onClick2).List()
	wg.Add(1)
	listen(onClick2, &wg)

	button.Emit(clickButtonEvent)
	button.Emit(mouseOverButtonEvent)

	wg.Wait()
}

func listen(c <-chan string, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		fmt.Println("Listener invoked ", <-c)
	}()
}
