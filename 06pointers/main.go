package main

import "fmt"

func main() {
	fmt.Println("Pointers in GoLang")

	// myVal := 10

	// var ptr1 *int
	// fmt.Println("Value of myPtr is: ", ptr1)

	myNumber := 10
	fmt.Println("Value of myNumber is: ", myNumber)
	var myPtr = &myNumber                             // & is used to get the address of a variable, can also write as myPtr := &myNumber, add type explicitly
	fmt.Println("Value of myPtr is: ", myPtr)         // reference to direct memory address
	fmt.Println("Value stored at myPtr is: ", *myPtr) // dereferencing the pointer, * is used to get the value stored at the address

	*myPtr *= 2 // updating the value stored at the address, thereby also updating the value of myNumber variable
	fmt.Println("New value of myNumber is: ", myNumber)
}
