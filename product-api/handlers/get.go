package handlers

import (
	"microservice-golang/data"
	"net/http"
)

// ListAll : handles GET requests and returns all current products
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] Get All Records")

	prods := data.GetProducts()

	err := data.ToJSON(prods, rw)
	if err != nil {
		p.l.Println("[ERROR] Serializing Product", err)
	}
}

// ListSingle : handles GET requests
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] Get Record ID", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] Fetching Product", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return

	default:
		p.l.Println("[ERROR] Fetching Product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		p.l.Println("[ERROR] Serializing Product", err)
	}
}
