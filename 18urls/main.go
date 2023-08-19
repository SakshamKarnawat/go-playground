package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://pokeapi.co:3000/api/v2/ability/?limit=20&offset=10"

func main() {
	fmt.Println("Handling URLs in Golang!")

	// Parsing URL
	parsedURL, _ := url.Parse(myurl)

	// Printing the parsed URL
	fmt.Println(parsedURL.Scheme)
	fmt.Println(parsedURL.Host)
	fmt.Println(parsedURL.Path)
	fmt.Println(parsedURL.Port())
	fmt.Println(parsedURL.RawQuery)

	// Getting the query parameters as a map
	queryParams := parsedURL.Query()

	// Printing the query parameters
	for key, value := range queryParams {
		fmt.Println("Key:", key, "Value:", value[0])
	}

	// The following shows why we need to use pointers when creating a URL struct

	partsOfURLValue := url.URL{
		Scheme:   "https",
		Host:     "pokeapi.co",
		Path:     "/api/v2/ability/",
		RawQuery: "limit=20&offset=10",
	}

	partsOfURLPointer := &url.URL{
		Scheme:   "https",
		Host:     "pokeapi.co",
		Path:     "/api/v2/ability/",
		RawQuery: "limit=20&offset=10",
	}

	fmt.Println("Before modifications:")
	fmt.Println("Value Struct:", partsOfURLValue.String())
	fmt.Println("Pointer Struct:", partsOfURLPointer.String())

	modifyURLValue(partsOfURLValue)
	modifyURLPointer(partsOfURLPointer)

	fmt.Println("\nAfter modifications:")
	fmt.Println("Value Struct:", partsOfURLValue.String())     // No change, since we passed a copy of the struct
	fmt.Println("Pointer Struct:", partsOfURLPointer.String()) // Changed, since we passed a pointer to the struct, so the original struct was modified
}

func modifyURLValue(u url.URL) {
	u.Scheme = "http"
	u.Host = "example.com"
	u.Path = "/modified-path"
	u.RawQuery = "param=modified"
}

func modifyURLPointer(u *url.URL) {
	u.Scheme = "http"
	u.Host = "example.com"
	u.Path = "/modified-path"
	u.RawQuery = "param=modified"
}
