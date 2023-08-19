package main

import "fmt"

func main() {
	fmt.Println("Defer in GoLang!")

	defer printABC() // last to be executed
	defer fmt.Println("\nWorld")
	defer myDefer()
	defer fmt.Println("Two")
	defer fmt.Println("One")
	fmt.Println("Hello")
}

func printABC() {
	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")
	// defer will be executed in LIFO order, i.e. C B A
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i, " ")
	}
}
