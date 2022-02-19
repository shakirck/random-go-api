// Package classification of Product API
//
// Documentation for PRoudcuct API

// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/josn
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/shakirck/go-micor/data"
)

// A list of prodcuts returs in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:response noContent
type prodctsNoContentWrapper struct {
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	//the id of the produc to be deleted form the database
	//in: path
	//required: true
	ID int `json:"id"`
}
type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := prod.FromJson(r.Body)
		if err != nil {
			http.Error(rw, "Unable to Marshal middleware ", http.StatusBadRequest)
			return
		}
		err = prod.Validate()
		if err != nil {
			http.Error(rw, fmt.Sprintf("Error validating the product %s", err), http.StatusBadRequest)
		}
		p.l.Printf("product ***** %#v", prod)
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
