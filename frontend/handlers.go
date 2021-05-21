package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"text/template"
)

func (f *FrontendServer) HomeHandler(w http.ResponseWriter, h *http.Request) {
	products := ListProducts()
	parsedTemplate, _ := template.ParseFiles("templates/home.html", "templates/header.html")
	err := parsedTemplate.Execute(w, products)
	if err != nil {
		log.Fatal("Error executing template:", err)
	}
}

func (f *FrontendServer) ProductHandler(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	product := ListProduct(productId)
	parsedTemplate, _ := template.ParseFiles("templates/product.html")
	err := parsedTemplate.Execute(w, product)
	if err != nil {
		log.Fatal("Error executing template:", err)
	}
}

