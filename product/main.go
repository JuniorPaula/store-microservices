package main

import (
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

func loadData() []byte {
	jsonFile, err := os.Open("products.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err.Error())
	}

	return data
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := loadData()
	w.Write([]byte(products))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", ListProducts)

	fmt.Println("server started on port [::8080]")
	http.ListenAndServe(":8080", r)
}
