package handlers

import (
	"net/http"

	"github.com/shakirck/go-micor/data"
)

// swagger:route GET /products products listProducts
// Returns a list of product
// Responses:
// 	200: productsResponse

// GetProducts return the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json  in get go ", http.StatusInternalServerError)
	}
}
