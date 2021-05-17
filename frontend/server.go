package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type FrontendServer struct {
	Router *mux.Router
}

func (f *FrontendServer) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, f.Router))
}

func (f *FrontendServer) InitializeRoutes() {
	f.Router.HandleFunc("/", f.HomeHandler).Methods("GET")
}
