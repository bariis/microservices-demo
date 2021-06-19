package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var products []Product
var product Product

func ListProducts() []Product {
	fmt.Println("burasi ListProducts()")
	resp, err := http.Get("http://localhost:5000/api/products")
	if err != nil {
		log.Fatalln("error in get request", err)
	}
	productJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error reading resp.body", err)
	}
	er := json.Unmarshal(productJson, &products)
	if er != nil {
		log.Fatal("error unmarshalling json", err)
	}
	return products
}

func ListProduct(productId string) Product {
	fmt.Println("burasi listproduct for single product and PRODUCT_ID:", productId)
	resp, err := http.Get("http://localhost:5000/api/products/1002")
	if err != nil {
		fmt.Println("Error fetching product with id: ", productId)
	}

	productJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response.body", err)
	}
	err = json.Unmarshal(productJson, &products)
	if err != nil {
		fmt.Println("Error reading response.body", err)
	}
	fmt.Println("PRODUCTUN SON HALI BURADA:", products)
	return products[0]
}
