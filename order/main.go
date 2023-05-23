package main

import (
	"fmt"
	"order/queue"
	"os"
)

type Product struct {
	UUID    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float64 `json:"price,string"`
}

type Order struct {
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	ProductID string `json:"product_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at,string"`
}

var productURL string

func init() {
	productURL = os.Getenv("PRODUCT_URL")
}

func main() {
	in := make(chan []byte)

	connection := queue.Connect()
	queue.StartConsumer(connection, in)

	for payload := range in {
		fmt.Println(string(payload))
	}
}
