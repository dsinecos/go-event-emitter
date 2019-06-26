package main

import (
	"fmt"
	"sync"

	EventEmitter "github.com/dsinecos/go-event-emitter/evtemtr"
)

func main() {
	button := EventEmitter.New()

	onClick1 := make(chan string)
	onClick2 := make(chan string)

	var wg sync.WaitGroup
	button.On("click", onClick1).List()
	wg.Add(1)
	listen(onClick1, &wg)
	button.On("click", onClick2).List()
	wg.Add(1)
	listen(onClick2, &wg)

	button.Emit("click")
	button.Emit("mouseover")

	wg.Wait()
}

func listen(c <-chan string, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		fmt.Println("Listener invoked ", <-c)
	}()
}
