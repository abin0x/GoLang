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
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
	w.Header().Set("Content-Type", "application/json") // Set content type to JSON
	if r.Method != http.MethodGet {                    // or use r.Method != "GET"  only GET requests allowed
		http.Error(w, "Method not allowed, please use GET", http.StatusMethodNotAllowed) //or use status code 400,syntax is  http.Error(w, "Method not allowed", 400)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

// createproductHandler handles POST requests to create a new product

func createproductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		http.Error(w, "Method not allowed, please use POST", http.StatusMethodNotAllowed)
		return
	}

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
