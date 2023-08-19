package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang!")

	// no inheritance in GoLang; no super / parent; no sub / child

	saksham := User{"Saksham", "test@abc.com", 21, true}
	fmt.Println(saksham)
	fmt.Printf("details are: %+v\n", saksham) // %+v gives the field names as well, %v gives only the values
	fmt.Printf("Name is %v and Email is %v\n", saksham.name, saksham.Email)
}

type User struct {
	name     string // small case name means private, capital means public
	Email    string
	Age      int
	Verified bool
}
