package main

import (
	"encoding/json"
	"fmt"
	"html/template"
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

	return products.Products
}

func ListProductsFromHTMLTemplate(w http.ResponseWriter, r *http.Request) {
	products := loadProducts()
	t := template.Must(template.ParseFiles("template/catalog.html"))
	t.Execute(w, products)
}

func ShowProductFromHTMLTemplate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response, err := http.Get(productURL + "/products/" + vars["id"])
	if err != nil {
		fmt.Printf("the HTTP request failed with error: %s\n", err)
	}

	data, _ := io.ReadAll(response.Body)

	var product Product
	json.Unmarshal(data, &product)

	t := template.Must(template.ParseFiles("template/view.html"))
	t.Execute(w, product)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", ListProductsFromHTMLTemplate)
	r.HandleFunc("/product/{id}", ShowProductFromHTMLTemplate)

	fmt.Print("catalog server started on port [::8081]\n")
	http.ListenAndServe(":8081", r)
}
