package http

import (
	"github.com/spacetronot-research-team/catalog-service/internal/service"
)

type CategoryController struct {
	productService service.Product
}

func NewCategoryController(productService service.Product) *CategoryController {
	return &CategoryController{
		productService: productService,
	}
}
