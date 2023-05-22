package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var productsURL string

func init() {
	productsURL = os.Getenv("PRODUCT_URL")
}

func displayCheckout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello checkout!"))
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/{id}", displayCheckout)

	fmt.Println("checkout server started on port 8082")
	http.ListenAndServe(":8082", r)
}
