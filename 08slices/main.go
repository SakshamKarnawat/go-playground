package main

import (
	"fmt"
	"sort"
)

func main() {
	println("Slices in Golang!")

	var nameList = []string{"John", "Doe", "Smith", "Jane", "Foo", "Bar"}

	nameList = append(nameList[1:4], "Bob") //  Doe, Smith, Jane, Bob because this is slicing from index 1 to 4 and then appending Bob
	fmt.Println(nameList)

	fmt.Println(nameList[1:3]) // Smith Jane, because this is slicing from index 1 to 3

	highScores := make([]int, 4) // array of 4 elements
	highScores[0] = 100
	highScores[1] = 120
	highScores[2] = 90
	highScores[3] = 80
	// highScores[4] = 104 // this will throw an error because the array is only 4 elements long

	highScores = append(highScores, 104, 153, 255) // memory is reallocated and the array is now 7 elements long
	fmt.Println("High Scores:", highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores) // sort the array
	fmt.Println("Sorted High Scores:", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// remove a value from slice based on index
	var languages = []string{"Go", "Python", "Java", "C#", "JavaScript", "Ruby", "PHP"}
	indexToRemove := 0
	languages = append(languages[:indexToRemove], languages[indexToRemove+1:]...)
	fmt.Println(languages)

	// remove a value from slice based on value
	var languages2 = []string{"Go", "Python", "Java", "C#", "JavaScript", "Ruby", "PHP"}
	valueToRemove := "C#"
	for index, value := range languages2 {
		if value == valueToRemove {
			languages2 = append(languages2[:index], languages2[index+1:]...)
			break
		}
	}
	fmt.Println(languages2)
}
