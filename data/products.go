package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Desc      string  `json:"desc"`
	Price     float32 `json:"price"`
	SKU       string  `json:"sku"`
	CreatedOn string  `json:"-"`
	UpdateOn  string  `json:"-"`
	DeleteOn  string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:        1,
		Name:      "Latte",
		Desc:      "Frothy milky coffee",
		Price:     2.45,
		SKU:       "abc123",
		CreatedOn: time.Now().UTC().String(),
		UpdateOn:  time.Now().UTC().String(),
	},
	&Product{
		ID:        2,
		Name:      "Espresso",
		Desc:      "Short and strong coffee without milk",
		Price:     1.99,
		SKU:       "mds",
		CreatedOn: time.Now().UTC().String(),
		UpdateOn:  time.Now().UTC().String(),
	},
}
