package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition in Golang!")

	wg := &sync.WaitGroup{} // using as pointer // right way to use wait groups
	mut := &sync.RWMutex{}  // using as pointer // right way to use mutex

	var score = []int{0}

	wg.Add(4)
	// IFFY (Immediately Invoked Function Expression) - These functions are called immediately after declaration
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Four R")
		m.RLock() // read lock
		fmt.Println("Score: ", score)
		m.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait() // wait for all the goroutines to finish, before exiting the main function

	fmt.Println("Final Score: ", score)
}

// go run --race . --> produces race condition error if lock and unlock are not used
