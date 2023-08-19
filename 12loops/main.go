package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in GoLang!")

	// if arrays are used instead of slices, then there might be initial values in the array which are not initialized
	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	fmt.Println("Days of the week are: ", days)

	// Lesser used loop syntax:
	/* for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	} */

	/* for i := range days {
		fmt.Println(days[i]) // i is not the value, but the index
	} */

	// For each loop syntax equivalent:
	/* for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	} */

	// While loop equivalent:
	val := 1
	for val <= 10 {

		if val == 3 {
			val++ // if we use continue without updating the value of val, then it will halt the program
			continue
		}

		if val == 5 {
			goto myLabel
		}

		if val == 8 {
			break
		}

		fmt.Println(val)
		val++
	}

myLabel:

	fmt.Println("This is my label")
}
