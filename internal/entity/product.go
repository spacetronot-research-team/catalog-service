package entity

import "errors"

var (
	ErrProductNotFound      = errors.New("product does not exist")
	ErrProductAlreadyExists = errors.New("product already exists")
)

type Product struct {
	ID       int64
	Name     string
	Category Category
	Price    float64
	Stock    int64
}
