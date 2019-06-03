package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

type Book struct {
	//backticks encodes the value of id into the key ID, here for json
	ID     int		`json:di`
	Title  string 	`json:title`
	Author string 	`json:author`
	Year   string 	`json:year`
}

var books []Book

func main() {

	books = append(books, Book{ID: 1, Title:"Golang Pointers", Author: "Mister Golang", Year: "2010"},
		Book{ID: 2, Title:"Golang routines", Author: "Mister Routines", Year: "2011"},
		Book{ID: 3, Title:"Golang routers", Author: "Mister Routers", Year: "2012"},
		Book{ID: 4, Title:"Golang  concurrency", Author: "Mister Currency", Year: "2013"},
		Book{ID: 5, Title:"Golang good parts", Author: "Mister Good", Year: "2014"},);

	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8888", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Get all books ...")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Get a single Book ...")
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])

	log.Println(reflect.TypeOf(i))

	for _, book := range books{
		if book.ID == i {
			fmt.Println("in if ###########")
			json.NewEncoder(w).Encode(books[i])
		}
	}
	log.Println(params)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Add a book ...")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Update a book ...")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove a book ...")
}
