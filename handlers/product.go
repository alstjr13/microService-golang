package handlers

import (
	"encoding/json"
	"log"
	"module/product-api/data"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()

	// encoding JSON
	// func Marshal(v interface{}) ([]byte, error)
	d, err := json.Marshal(lp)

	if err != nil {
		http.Error(rw, "Unable to Marshal JSON", http.StatusInternalServerError)
	}

	rw.Write(d)
}
