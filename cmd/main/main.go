package main

import (
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/* basically in this file we will create the server, which will also define our local
host. Here, we will also tell Golang where our routes are located. */

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
