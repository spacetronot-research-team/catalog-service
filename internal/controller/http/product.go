package http

import (
	"github.com/spacetronot-research-team/catalog-service/internal/service"
)

type ProductController struct {
	productService service.Product
}

func NewProductController(productService service.Product) *ProductController {
	return &ProductController{
		productService: productService,
	}
}
