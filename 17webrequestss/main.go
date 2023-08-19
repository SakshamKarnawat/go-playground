package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://catfact.ninja/fact"

func main() {
	fmt.Println("Web requests in Golang!")

	// http get request:
	resp, err := http.Get(url)

	checkNilError(err)

	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close() // this is important to close the response body

	fmt.Printf("Response status: %s\n", resp.Status)

	bytes, err := io.ReadAll(resp.Body)

	checkNilError(err)

	body := string(bytes)

	fmt.Println("Response body: ", body)

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
