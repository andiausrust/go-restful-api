package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"github.com/andiausrust/books.list/driver"
	"github.com/andiausrust/books.list/models"
	"github.com/andiausrust/books.list/controllers"
)

var books []models.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {

	db = driver.ConnectDB()
	router := mux.NewRouter()

	controller := controllers.Controller{}

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")


	log.Println("Server is running on Port: 8888")
	driver.LogFatal(http.ListenAndServe(":8888", router))
}


