package main

import "fmt"

func main() {
	fmt.Println("Maps in GoLang!")

	languages := make(map[string]string) // key: string, value: string
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"
	languages["GO"] = "GOOOOOO" // overwrites the value of GO key

	fmt.Println("Languages: ", languages)
	fmt.Println("JS short for: ", languages["JS"])

	delete(languages, "PY") // deletes the key-value pair "PY"
	fmt.Println("Languages: ", languages)

	// Iterating over maps
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
