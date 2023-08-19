package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to user input!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok || err ok
	input, err := reader.ReadString('\n') // this will read the input until it finds a new line character
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again", err)
		return
	}
	// fmt.Print("Thanks for rating ", input)
	// fmt.Printf("Type of the rating: %T \n", input)

	// convert string to float and remove the trailing new line character
	var numRating, err2 = strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err2 != nil {
		fmt.Println("An error occurred while reading input. Please try again", err2)
		return
	} else {
		fmt.Println("Thanks for rating ", numRating+1)
	}
}
