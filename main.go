package main

import (
	"net/http"

	"github.com/goddeuce1/highload-hw-2/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/main", handlers.HandlerMain).Methods("GET")
	router.HandleFunc("/shop", handlers.HandlerShop).Methods("GET")
	router.HandleFunc("/blog", handlers.HandlerBlog).Methods("GET")

	http.ListenAndServe(":8080", router)

}
