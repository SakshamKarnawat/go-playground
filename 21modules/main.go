package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Modules in Golang!")

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")

	// http.Handle("/", r)

	log.Fatal(
		http.ListenAndServe(":8080", r),
	)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello World!")
	w.Write([]byte("<h1>Hello World!</h1>"))
	fmt.Println("lmao")
	// log to browser console:

}

// TODO: go mod verify --- can't resolve error for this
