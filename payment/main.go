package main

import (
	"encoding/json"
	"fmt"
	"payment/queue"
	"time"

	"github.com/streadway/amqp"
)

type Order struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	ProductID string    `json:"product_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at,string"`
}

func notifyPaymentOrder(order Order, ch *amqp.Channel) {
	json, _ := json.Marshal(order)
	queue.Notify(json, "payment_ex", "", ch)

	fmt.Println(string(json))
}

func main() {
	in := make(chan []byte)

	connection := queue.Connect()
	queue.StartConsumer("order_queue", connection, in)

	var order Order
	for payload := range in {
		json.Unmarshal(payload, &order)
		order.Status = "approved"

		notifyPaymentOrder(order, connection)
	}

}
