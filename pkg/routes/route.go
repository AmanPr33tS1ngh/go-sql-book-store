package routes

import (
	"github.com/AmanPr33tS1ngh/go-sql-book-store.git/pkg/controllers"
	"github.com/gorilla/mux"
)

var BookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}/", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}/", controllers.DeleteBook).Methods("DELETE")
}