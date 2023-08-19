package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to my time in Go!")
	dateFormat := "02 January 03:04:05PM 2006 MST"

	currentTime := time.Now()
	fmt.Println("Current time is: ", currentTime)

	fmt.Println(currentTime.Format(dateFormat))

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.Local)
	fmt.Println("Created date is: ", createdDate)
	fmt.Println("Formatted created date is: ", createdDate.Format(dateFormat))

}
