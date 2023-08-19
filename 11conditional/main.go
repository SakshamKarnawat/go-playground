package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Conditional Statements in GoLang")

	// If else statement
	fmt.Println("If else statement")

	var result string
	num := 23

	if num < 23 {
		result = "Number is less than 23"
	} else if num > 23 {
		result = "Number is greater than 23"
	} else {
		result = "Number is equal to 23"
	}

	fmt.Println("Result is: ", result)

	if newNum := 5; newNum < 23 { // newNum is only available in if block
		result = "New Number is less than 23"
	}

	// Switch statement
	fmt.Println("Switch statement")

	diceNum := rand.Intn(6) + 1

	switch diceNum {
	case 1:
		fmt.Println("Dice value is 1")
		fallthrough // fallthrough is used to execute next case
	case 2:
		fmt.Println("Dice value is 2")
		fallthrough
	case 3:
		fmt.Println("Dice value is 3")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4")
		fallthrough
	case 5:
		fmt.Println("Dice value is 5")
		fallthrough
	case 6:
		fmt.Println("Dice value is 6")
		fallthrough
	default:
		fmt.Println("Dice value is unknown")
	}
}
