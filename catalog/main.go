package main

import "os"

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

func main() {

}
