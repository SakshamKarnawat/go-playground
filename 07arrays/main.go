package main

import (
	"fmt"
)

func main() {
	println("Arrays in GoLang")

	var nameList [4]string

	nameList[0] = "John"
	nameList[1] = "Paul"
	// nameList[2] = "George"
	nameList[3] = "Ringo"
	// element 2 is not assigned a value, so it will be assigned the default value of the type, which is an empty string in this case

	fmt.Println("nameList is: ", nameList)
	fmt.Println("nameList length is: ", len(nameList))

	var numList = [10]int{1, 2, 3, 4}
	fmt.Println("numList is: ", numList)
	fmt.Println("numList length is: ", len(numList)) // length is 10, even though only 4 elements are assigned values
}
