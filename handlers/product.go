package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/shakirck/go-micor/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		p.l.Println(r.URL.Path)
		regx := regexp.MustCompile(`/([0-9]+)`)
		g := regx.FindAllSubmatch([]byte(r.URL.Path), -1)

		if len(g) != 1 {
			http.Error(rw, "Invalid URI no id/ or multiple  ", http.StatusBadRequest)
		}
		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URI ", http.StatusBadRequest)
		}
		idString := g[0][1]
		id, err := strconv.Atoi(string(idString))
		if err != nil {
			http.Error(rw, " Invalid URI", http.StatusBadRequest)
		}
		p.l.Println(id)
		p.updateProducts(id, rw, r)
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json ", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("adding product POST")

	prod := &data.Product{}

	err := prod.FromJson(r.Body)

	if err != nil {
		http.Error(rw, "Unable to Marshal ", http.StatusBadRequest)
	}
	p.l.Printf("%#v", prod)
	data.AddProduct(prod)

}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("adding product POST")

	prod := &data.Product{}

	err := prod.FromJson(r.Body)

	if err != nil {
		http.Error(rw, "Unable to Marshal ", http.StatusBadRequest)
	}
	p.l.Printf("%#v", prod)
	err = data.UpdateProduct(id, prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product Not found ", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "product not found ", http.StatusInternalServerError)
	}
}
