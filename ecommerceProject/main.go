package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// helloHandler handles requests to the /hello endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!!")
}

// aboutHandler handles requests to the /about endpoint
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About us page")
}

// Product struct represents a product in the ecommerce application
type Product struct {
	ID          int `json:"id"`
	Title       string
	Description string
	Price       float64
	ImgUrl      string
}

// global slice to store products
var productList []Product

// getProductsHandler handles GET requests to the /products endpoint
func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { // or use r.Method != "GET"  only GET requests allowed
		http.Error(w, "Method not allowed, please use GET", http.StatusMethodNotAllowed) //or use status code 400,syntax is  http.Error(w, "Method not allowed", 400)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

// main function sets up the HTTP server and routes
func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/products", getProductsHandler)
	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
func init() {
	// Initialize the product list with some sample products
	prd1 := Product{
		ID:          1,
		Title:       "Laptop",
		Description: "A high-performance laptop",
		Price:       999.99,
		ImgUrl:      "http://example.com/laptop.jpg",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Smartphone",
		Description: "A latest model smartphone",
		Price:       699.99,
		ImgUrl:      "http://example.com/smartphone.jpg",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Headphones",
		Description: "Noise-cancelling headphones",
		Price:       199.99,
		ImgUrl:      "http://example.com/headphones.jpg",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Smartwatch",
		Description: "A smartwatch with various features",
		Price:       299.99,
		ImgUrl:      "http://example.com/smartwatch.jpg",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Tablet",
		Description: "A lightweight tablet",
		Price:       399.99,
		ImgUrl:      "http://example.com/tablet.jpg",
	}
	prd6 := Product{
		ID:          6,
		Title:       "Camera",
		Description: "A digital camera",
		Price:       499.99,
		ImgUrl:      "http://example.com/camera.jpg",
	}
	// productList = []Product{prd1, prd2, prd3, prd4, prd5, prd6}
	productList = append(productList, prd1, prd2, prd3, prd4, prd5, prd6)
}
