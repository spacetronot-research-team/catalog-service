package model

type Product struct {
	ID       int64
	Category Category
	Name     string
	Stock    int64
	Price    Price
}
