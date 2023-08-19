package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions in GoLang!")
	a := greeter
	a()

	greeterTwo := func() {
		fmt.Println("This is an anonymous function in GoLang!")
	}
	greeterTwo()

	func() {
		fmt.Println("This is an IIFE (Immediately Invoked Function Expression) in GoLang!")
	}()

	result := add(2, 3)
	fmt.Println("The result of add is:", result)

	result, status := addAll(10, 20, 30, 40, 50)
	fmt.Println("The result of addAll is:", result)
	fmt.Println("Sum is more than 100?", status)
}

func greeter() {
	fmt.Println("Hello World!")
}

func add(a, b int) int {
	return a + b
}

// Variadic Function (Variable number of arguments)
func addAll(nums ...int) (int, bool) { // multiple return types using paranthesis
	fmt.Println("This is a variadic function in GoLang!")
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum, sum > 100
}
