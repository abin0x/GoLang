package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createproductHandler(w http.ResponseWriter, r *http.Request) {

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
