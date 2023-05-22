package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

type Product struct {
	UUID    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float64 `json:"price,string"`
}

type Products struct {
	Products []Product
}

type Order struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	ProductID string `json:"product_id"`
}

var productsURL string

func init() {
	productsURL = os.Getenv("PRODUCT_URL")
}

func displayCheckout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response, err := http.Get(productsURL + "/products/" + vars["id"])
	if err != nil {
		fmt.Println("the HTTP request falied with error: ", err)
	}

	data, _ := io.ReadAll(response.Body)

	var product Product
	json.Unmarshal(data, &product)

	t := template.Must(template.ParseFiles("template/checkout.html"))
	t.Execute(w, product)

}

func finish(w http.ResponseWriter, r *http.Request) {
	var order Order
	order.Name = r.FormValue("name")
	order.Email = r.FormValue("email")
	order.Phone = r.FormValue("phone")
	order.ProductID = r.FormValue("product_id")

	data, _ := json.Marshal(order)
	fmt.Println(string(data))

	w.Write([]byte("Procceced!!!"))
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/finish", finish)
	r.HandleFunc("/{id}", displayCheckout)

	fmt.Println("checkout server started on port 8082")
	http.ListenAndServe(":8082", r)
}
