package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shakirck/go-micor/data"
)

// swagger:route PUT / products updateProducts
// updates a product
//  consumes:
//  - application/json
// Responses:
//  200: productsNoResponse

func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "invalid Id ", http.StatusForbidden)
	}
	p.l.Println("adding product PUT", id)

	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	err = data.UpdateProduct(id, prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product Not found 1", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "product not found 2", http.StatusInternalServerError)
	}
}
