package main

type Product struct {
	ID          int     `json:"product_id"`
	Name        string  `json:"product_name"`
	Description string  `json:"product_description"`
	Price       float32 `json:"product_price"`
}
