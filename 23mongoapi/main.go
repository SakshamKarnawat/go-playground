package main

import (
	"fmt"
	"log"
	"mongoapi/router"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API in Golang!")

	fmt.Println("Starting server")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server started on port 8080...")
}
