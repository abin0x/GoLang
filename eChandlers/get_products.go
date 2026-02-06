package eChandlers

import "net/http"

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	senData(w, productList, 200)
}
