package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type product struct {
	Id       string
	Name     string
	Quantity int
	Price    float64
}

var Products []product

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcom to the home page")
	log.Println("Endpoint hit :homepage")
}

// ............All product Id..................
func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Welcom to the product page")
	json.NewEncoder(w).Encode(Products)
	log.Println("product page", Products)
}

//.................. Specific product......................

func getProduct(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Welcom to the product page")

	// log.Println(r.URL.Path)
	// key := r.URL.Path[len("/products/"):]

	vars := mux.Vars(r)
	key := vars["id"]

	for _, product := range Products {
		if string(product.Id) == key {
			json.NewEncoder(w).Encode(product)
		}
	}

}

func handleRequests() {
	// http.HandleFunc("/", homepage)
	// http.HandleFunc("/products", returnAllProducts)
	// http.HandleFunc("/products/", getProduct)

	MyRouter := mux.NewRouter().StrictSlash(true)
	MyRouter.HandleFunc("/", homepage)
	MyRouter.HandleFunc("/products", returnAllProducts)
	MyRouter.HandleFunc("/products/{id}", getProduct)

	//  running on port number 10000
	// http.ListenAndServe("localhost:10000", nil)
	http.ListenAndServe("localhost:10000", MyRouter)
}

func main() {
	Products = []product{
		product{Id: "1", Name: "Chair", Quantity: 30, Price: 100.00},
		product{Id: "2", Name: "Desk", Quantity: 39, Price: 300.00},
	}

	handleRequests()

}
