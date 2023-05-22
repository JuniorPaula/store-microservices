package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

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

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/{id}", displayCheckout)

	fmt.Println("checkout server started on port 8082")
	http.ListenAndServe(":8082", r)
}
