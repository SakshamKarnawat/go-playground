package main

import (
	"fmt"
	"net/http"
	"sync"
)

// * Goroutines are lightweight threads managed by the Go runtime

var signals = []int{}

var wg sync.WaitGroup // wait group is used to wait for a collection of goroutines to finish // used as pointer in actual code
var mut sync.Mutex    // mutex is used to lock a variable, so that only one goroutine can access it at a time // pointer

func main() {
	// go greeter("Hello")
	// greeter("World")

	websiteList := []string{
		"https://github.com",
		"https://go.dev",
		"https://google.com",
		"https://facebook.com",
		"https://flutter.dev",
	}

	for _, website := range websiteList {
		go getStatusCode(website)
		wg.Add(1) // add 1 to the wait group, so that the main function waits for the goroutine to finish
	}

	wg.Wait() // wait for all the goroutines to finish, before exiting the main function

	fmt.Println("Signals: ", signals)
}

// * not an efficient way to do handle goroutines
// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done() // remove 1 from the wait group, when the goroutine is finished

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Oops in endpoint")
	} else {
		mut.Lock() // lock the variable, so that only one goroutine can access it at a time
		signals = append(signals, res.StatusCode)
		mut.Unlock() // unlock the variable, so that other goroutines can access it
		fmt.Printf("%d Status code for %s\n", res.StatusCode, endpoint)
	}
}
