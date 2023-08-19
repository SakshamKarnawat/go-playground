package main

import "fmt"

// public variable - starts with capital letter - can be accessed outside the package if imported (example: LoginToken)
const LoginToken string = "asdasdasdasdasd"

// private variable - starts with small letter - can be accessed only inside the package (example: password)
var password = "pass"

func main() {
	// Declare a variable and initialize it
	name := "saksham"
	fmt.Println(name)
	fmt.Printf("name is of type: %T \n", name)

	// Declare a bool variable and initialize it
	var isTrue bool = true
	fmt.Println(password)
	fmt.Printf("isTrue is of type: %T \n", isTrue)

}
