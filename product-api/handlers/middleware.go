package handlers

import (
	"microservice-golang/data"
	"net/http"
)

func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")

		prod := &data.Product{}

		err := data.ToJSON(prod, r.Body)

		if err != nil {
			p.l.Error()

		}
	})
}
