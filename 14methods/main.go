package main

import "fmt"

func main() {
	fmt.Println("Methods in GoLang!")

	saksham := User{"Saksham", "test@abc.com", 21, true}
	fmt.Println(saksham)
	fmt.Printf("details are: %+v\n", saksham)
	fmt.Printf("Name is %v and Email is %v\n", saksham.Name, saksham.Email)

	saksham.GetVerifiedStatus() // calling method

	saksham.SetName("Saksham 2")
	fmt.Printf("Name is %v\n", saksham.GetName())
	// original name is not changed, because we are not passing pointer, so new copy is created
}

type User struct {
	Name     string
	Email    string
	Age      int
	Verified bool
}

// defining method for struct User
func (u User) GetVerifiedStatus() {
	fmt.Println("Is user verified: ", u.Verified)
}

func (u User) GetName() string {
	return u.Name
}

func (u User) SetName(name string) {
	u.Name = name
	fmt.Println("Name changed to: ", u.Name)
}
