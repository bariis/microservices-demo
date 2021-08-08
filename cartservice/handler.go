package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func sendResponse(w http.ResponseWriter, data []byte, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func (s *server) handleAddCart(w http.ResponseWriter, r *http.Request) {
	type product struct {
		Id       string `json:"id"`
		Quantity int    `json:"quantity"`
	}
	var p product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := r.Header.Get("Authorization")

	if err := AddItem(s.redisClient, p.Id, p.Quantity, token); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(map[string]string{"message": "item added to the cart successfully", "productId": p.Id, "productQuantity": strconv.Itoa(p.Quantity)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendResponse(w, jsonResponse, http.StatusOK)
}

func (s *server) handleGetCart(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	productItem, err := GetCart(s.redisClient, token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productJson, err := json.Marshal(productItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendResponse(w, productJson, http.StatusOK)
}

func (s *server) handleEmptyCart(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	if err := EmptyCart(s.redisClient, token); err != nil {
		http.Error(w, err.Error(), http.StatusOK)
		return
	}

	jsonResponse, err := json.Marshal(map[string]string{"message": "cart is deleted successfully"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendResponse(w, jsonResponse, http.StatusOK)
}
