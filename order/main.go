package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"order/queue"
	"os"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

type Product struct {
	UUID    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float64 `json:"price,string"`
}

type Order struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	ProductID string    `json:"product_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at,string"`
}

var productURL string

func init() {
	productURL = os.Getenv("PRODUCT_URL")
}

func createOrder(payload []byte) {
	var order Order
	json.Unmarshal(payload, &order)

	id, _ := uuid.NewV4()
	order.UUID = id.String()
	order.Status = "pending"
	order.CreatedAt = time.Now()

	saveOrder(order)
}

func saveOrder(order Order) {

}

func getProductByID(id string) Product {
	response, err := http.Get(productURL + "/products/" + id)
	if err != nil {
		fmt.Printf("The HTTP request failed with error: %s\n", err)
	}

	data, _ := io.ReadAll(response.Body)

	var product Product
	json.Unmarshal(data, &product)
	return product
}

func main() {
	in := make(chan []byte)

	connection := queue.Connect()
	queue.StartConsumer(connection, in)

	for payload := range in {
		createOrder(payload)
		fmt.Println(string(payload))
	}
}
