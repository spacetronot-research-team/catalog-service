package service

import (
	"github.com/spacetronot-research-team/catalog-service/internal/repository"
)

type Product interface {
	Create()
	GetList()
	GetDetails()
	Update()
	Delete()
}

type productService struct {
	productRepository repository.Product
}

func NewProductService(productRepository repository.Product) Product {
	return &productService{
		productRepository: productRepository,
	}
}

// Create implements Product.
func (*productService) Create() {
	panic("unimplemented")
}

// Delete implements Product.
func (*productService) Delete() {
	panic("unimplemented")
}

// GetDetails implements Product.
func (*productService) GetDetails() {
	panic("unimplemented")
}

// GetList implements Product.
func (*productService) GetList() {
	panic("unimplemented")
}

// Update implements Product.
func (*productService) Update() {
	panic("unimplemented")
}
