package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"order/db"
	"order/queue"
	"os"
	"time"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/streadway/amqp"
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

func createOrder(payload []byte) Order {
	var order Order
	json.Unmarshal(payload, &order)

	id, _ := uuid.NewV4()
	order.UUID = id.String()
	order.Status = "pending"
	order.CreatedAt = time.Now()

	saveOrder(order)
	return order
}

func saveOrder(order Order) {
	json, _ := json.Marshal(order)
	connection := db.Connect()

	err := connection.Set(order.UUID, string(json), 0).Err()
	if err != nil {
		panic(err.Error())
	}
}

func notifyOrderCreated(order Order, ch *amqp.Channel) {
	json, _ := json.Marshal(order)
	queue.Notify(json, "order_ex", "", ch)
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
	var param string
	flag.StringVar(&param, "opt", "", "Usage")
	flag.Parse()

	in := make(chan []byte)
	connection := queue.Connect()

	switch param {
	case "checkout":
		queue.StartConsumer("checkout_queue", connection, in)

		for payload := range in {
			notifyOrderCreated(createOrder(payload), connection)
			fmt.Println("#Checkout;", string(payload))
		}
	case "payment":
		queue.StartConsumer("payment_queue", connection, in)

		var order Order
		for payload := range in {
			json.Unmarshal(payload, &order)
			saveOrder(order)

			fmt.Println("#Payment;", string(payload))
		}
	}

}
