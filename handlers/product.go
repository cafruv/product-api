package handlers

import (
	"log"
	"net/http"

	"github.com/cafruv/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) SerceHTTP(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
  d, err : = json.Marshal(lp)
  if err!= nil {
    http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
  }

  rw.Write(d)
}
