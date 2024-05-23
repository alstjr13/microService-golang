package handlers

import (
	"microservice-golang/data"
	"net/http"
)

func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("[DEBUG] Updating record ID", prod.ID)

	err := data.UpdateProduct(prod)

	// If the error is the ErrProductNotFound error
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] Product Not Foound", err)

		// Not found Error
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product Not Found in the Database"}, rw)
		return
	}

	// Write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}
