package service

import (
	"context"

	"github.com/spacetronot-research-team/catalog-service/internal/model"
	"github.com/spacetronot-research-team/catalog-service/internal/repository"
)

type Product interface {
	Create()
	GetList()
	GetDetail(ctx context.Context, id int) (product model.Product, err error)
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
func (ps *productService) GetDetail(ctx context.Context, id int) (product model.Product, err error) {
	product, err = ps.productRepository.GetDetail(ctx, id)

	return
}

// GetList implements Product.
func (*productService) GetList() {
	panic("unimplemented")
}

// Update implements Product.
func (*productService) Update() {
	panic("unimplemented")
}
