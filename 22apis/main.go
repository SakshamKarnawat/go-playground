package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for Book - file
type Book struct {
	BookId    string  `json:"id"`
	BookName  string  `json:"name"`
	BookPrice int     `json:"price"`
	Author    *Author `json:"author"` // not using pointer will cause issues when updating data, since it will be a copy of the data
}

// Model for Author - file
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake database:
var books []Book

// middleware, helper - file
func (b *Book) isEmpty() bool {
	return b.BookName == ""
}

func main() {
	fmt.Println("APIs in Golang!")

	r := mux.NewRouter()

	// seeding:
	books = append(
		books,
		Book{BookId: "1", BookName: "Book 1", BookPrice: 199, Author: &Author{Fullname: "Author 1", Website: "https://author1.com"}},
		Book{BookId: "2", BookName: "Book 2", BookPrice: 299, Author: &Author{Fullname: "Author 2", Website: "https://author2.com"}},
		Book{BookId: "3", BookName: "Book 3", BookPrice: 399, Author: &Author{Fullname: "Author 3", Website: "https://author3.com"}},
	)

	// routes:
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/book", createBook).Methods("POST")
	r.HandleFunc("/book/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")

	// listening to port 8080:
	log.Fatal(http.ListenAndServe(":8080", r))

	// srv := &http.Server{
	// 	Handler: r,
	// 	Addr:    "127.0.0.1:8080",
	// 	// Good practice: enforce timeouts for servers you create!
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }
	// log.Fatal(srv.ListenAndServe())

}

// controllers - file
func getBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside getBooks controller")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside getBook controller")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, book := range books {
		if book.BookId == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode("Book not found for the given id - " + params["id"])
}

func createBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside createBook controller")
	w.Header().Set("Content-Type", "application/json")

	// if body is empty:
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send book details")
		return
	}

	// {}

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	if book.isEmpty() {
		json.NewEncoder(w).Encode("No book details inside JSON")
		return
	}

	// if book name is same as existing book name, then return saying book already exists:
	for _, b := range books {
		if b.BookName == book.BookName {
			json.NewEncoder(w).Encode("Book already exists")
			return
		}
	}

	// genrerate a unique id and convert to string
	book.BookId = strconv.Itoa(rand.Intn(1000))

	books = append(books, book)
	json.NewEncoder(w).Encode(book)
	return
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside updateBook controller")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	oldBooks := books // copy of the books, since we are going to update the books, and we shouldn't modify the slice while iterating

	for index, book := range oldBooks {
		if book.BookId == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var newBook Book // can't name it book, since we are using book in the for loop
			_ = json.NewDecoder(r.Body).Decode(&newBook)
			newBook.BookId = params["id"]
			books = append(books, newBook)
			json.NewEncoder(w).Encode(books)
			return
		}
	}

	// TODO: send a response if book not found
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside deleteBook controller")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	oldBooks := books // copy of the books, since we are going to update the books, and we shouldn't modify the slice while iterating

	for index, book := range oldBooks {
		if book.BookId == params["id"] {
			books = append(books[:index], books[index+1:]...)
			json.NewEncoder(w).Encode(books)
			return
		}
	}

	// TODO: send a response if book not found
}

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Welcome to the home page! </h1>")

}
