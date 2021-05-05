package app

import (
	"database/sql"
	"errors"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *Product) getPrduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Product) updateProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Product) deleteProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Product) createProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getProducts(db *sql.DB, start, count int) ([]Product, error) {
	return nil, errors.New("Not implemented")
}
