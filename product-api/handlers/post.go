package handlers

import (
	"microservice-golang/data"
	"net/http"
)

func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("[DEBUG] Inserting Product: %#v\n", prod)
	data.AddProduct(prod)
}
