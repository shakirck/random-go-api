package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shakirck/go-micor/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Delete a Product from database
// Responses:
// 	200: noContent

// Delete  a prpduct  from the data store
func (p Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	log.Println("deleteing *****")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "invalid Id ", http.StatusForbidden)
	}
	p.l.Println("deleteing  product [DELETE]", id)

	err = data.DeleteProduct(id)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product Not found 1", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "product not found 2", http.StatusInternalServerError)
	}
}
