package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Working with txt files in Golang!")

	content := "This is the content of the file"

	// open file
	file, err := os.Create("./file.txt")

	checkNilError(err)

	// write to file
	length, err := file.WriteString(content)

	checkNilError(err)

	fmt.Printf("All done with file of %v characters\n", length)

	defer file.Close() // defer because we want to close the file after we are done with all the operations

	// read file
	readFile("./file.txt")
}

func readFile(filename string) {
	readBytes, err := os.ReadFile(filename) // bytes format
	checkNilError(err)
	fmt.Println(string(readBytes))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
