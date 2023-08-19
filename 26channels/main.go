package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang!")

	// myCh := make(chan int) // unbuffered channel, can only hold 1 value at a time, needs as many listeners as value sent
	myCh := make(chan int, 2) // buffered channel with capacity 2, can hold 2 values at a time
	wg := &sync.WaitGroup{}

	// * Gives error when used without goroutines
	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)

	// Receive only channel
	go func(ch <-chan int, wg *sync.WaitGroup) { // <-chan int is a receive only channel, by default channels are bidirectional
		// close(ch) // * Gives error when used with receive only channel

		// Receivng value from channel
		val, isOpen := <-ch

		fmt.Println(isOpen)
		fmt.Println(val)

		wg.Done()
	}(myCh, wg)

	// Send only channel
	go func(ch chan<- int, wg *sync.WaitGroup) { // chan<- int is a send only channel, by default channels are bidirectional
		// Putting value in channel
		ch <- 0
		close(ch)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
