package main

import "github.com/gorilla/mux"

func main() {
	s := server{}
	s.router = mux.NewRouter()
	s.router.HandleFunc("/api/v1/checkout", s.HandleCheckout).Methods("GET")
	s.Run()
}
