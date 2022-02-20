package handlers

import (
	"net/http"

	"github.com/shakirck/go-micor/data"
)

// swagger:route POST / products addProducts
// Add a product to the store
// Responses:
// 	200: productsNoResponse

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("adding product POST")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("%#v", prod)
	data.AddProduct(prod)

}
