package main

import (
	"log"
	"net/http"

	"github.com/AmanPr33tS1ngh/go-sql-book-store.git/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main()  {
	r := mux.NewRouter()
	routes.BookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}