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
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

// global slice to store products
var productList []Product

// getProductsHandler handles GET requests to the /products endpoint
func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	/*
		if r.Method != "OPTIONS" {
			w.WriteHeader(200)
			return
		}

	*/

	// handlePreflightRequest(w, r)
	if r.Method != http.MethodGet { // or use r.Method != "GET"  only GET requests allowed
		http.Error(w, "Method not allowed, please use GET", http.StatusMethodNotAllowed) //or use status code 400,syntax is  http.Error(w, "Method not allowed", 400)
		return
	}

	senData(w, productList, 200)
}

// createproductHandler handles POST requests to create a new product

func createproductHandler(w http.ResponseWriter, r *http.Request) {
	handleCors(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	/*
		handlePreflightRequest(w, r)
		if r.Method != "POST" {
			http.Error(w, "Method not allowed, please use POST", http.StatusMethodNotAllowed)
			return
		}
	*/

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)
	senData(w, newProduct, http.StatusCreated)

}

// handleCors sets the necessary CORS headers
func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

// handlePreflightRequest handles OPTIONS requests for CORS preflight
func handlePreflightRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}
}

// senData sends a JSON response with the given data and status code
func senData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

// main function sets up the HTTP server and routes
func main() {

	mux := http.NewServeMux()
	mux.Handle("GET /hello", http.HandlerFunc(helloHandler))
	mux.Handle("GET /about", http.HandlerFunc(aboutHandler))
	mux.Handle("GET /products", http.HandlerFunc(getProductsHandler))
	mux.Handle("OPTIONS /products", http.HandlerFunc(getProductsHandler))
	mux.Handle("POST /createproduct", http.HandlerFunc(createproductHandler))
	mux.Handle("OPTIONS /createproduct", http.HandlerFunc(createproductHandler))
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
		ImgUrl:      "https://techterms.com/img/xl/laptop_586.png",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Smartphone",
		Description: "A latest model smartphone",
		Price:       699.99,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTTYAx3hqK3b3i-JGJlw7nhqqrccJtkMnNv8Q&s",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Headphones",
		Description: "Noise-cancelling headphones",
		Price:       199.99,
		ImgUrl:      "https://cdn.mos.cms.futurecdn.net/HMCWShKerkfeNQmYYhE3p7.jpg",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Smartwatch",
		Description: "A smartwatch with various features",
		Price:       299.99,
		ImgUrl:      "https://istarmax.com/wp-content/uploads/2024/04/Starmax-Product-Range-Summer-2024-2.png",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Tablet",
		Description: "A lightweight tablet",
		Price:       399.99,
		ImgUrl:      "https://p2-ofp.static.pub/fes/cms/2021/10/28/juqs65pgl1gh3dysi7yv1tnvtsiqva364946.png",
	}
	prd6 := Product{
		ID:          6,
		Title:       "Camera",
		Description: "A digital camera",
		Price:       499.99,
		ImgUrl:      "https://www.gearpatrol.com/wp-content/uploads/sites/2/2023/09/best-vintage-film-cameras-refresh-lead-650de2dd54d04-jpg.webp",
	}
	// productList = []Product{prd1, prd2, prd3, prd4, prd5, prd6}
	productList = append(productList, prd1, prd2, prd3, prd4, prd5, prd6)
}
