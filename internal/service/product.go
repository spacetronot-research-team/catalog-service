package service

import (
	"github.com/spacetronot-research-team/catalog-service/internal/dto"
	"github.com/spacetronot-research-team/catalog-service/internal/repository"
)

type Product interface {
	Create()
	GetLists()
	GetDetails() dto.CreateCategoryResponse
	Update()
	Delete()
}

type productService struct {
	productRepository repository.Product
}

func NewProductService(productRepository repository.Product) Product {
	return &productService{productRepository: productRepository}
}

func (*productService) Create() {
	panic("unimplemented")
}

func (*productService) Delete() {
	panic("unimplemented")
}

func (*productService) GetDetails() dto.CreateCategoryResponse {
	panic("unimplemented")
}

func (*productService) GetLists() {
	panic("unimplemented")
}

func (*productService) Update() {
	panic("unimplemented")
}
