package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

// NewHello creates a new hello handler with the given logger l
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	h.l.Println("Hello world")

	// read body
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oooops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", d)
}
