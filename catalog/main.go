package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Product struct {
	UUID    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float64 `json:"price,string"`
}

type Products struct {
	Products []Product
}

var productURL string

func init() {
	productURL = os.Getenv("PRODUCT_URL")
}

func loadProducts() []Product {
	response, err := http.Get(productURL + "/products")
	if err != nil {
		fmt.Println("[ERROR]", err)
	}
	data, _ := io.ReadAll(response.Body)

	var products Products
	json.Unmarshal(data, &products)

	fmt.Println(string(data))

	return products.Products
}

func main() {
	loadProducts()
}
