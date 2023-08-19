package main

import (
	crypto "crypto/rand"
	"fmt"
	"math/big"
	math "math/rand"
)

func main() {
	fmt.Println("Welcome to maths in Go!")

	// var numOne int = 1
	// var numTwo float64 = 2.0

	// fmt.Printf("Addition: %T", int(float64(numOne)+numTwo))

	//random number using math/rand
	fmt.Println("Random number using math/rand: ", math.Intn(10))

	//random number using crypto/rand
	randomNum, _ := crypto.Int(crypto.Reader, big.NewInt(10))
	fmt.Println("Random number using crypto/rand: ", randomNum)
}
