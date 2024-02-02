package router

import (
	"github.com/spacetronot-research-team/catalog-service/internal/controller/http"
	"github.com/spacetronot-research-team/catalog-service/internal/repository"
	"github.com/spacetronot-research-team/catalog-service/internal/service"
	"gorm.io/gorm"
)

func getProductController(db *gorm.DB) *http.ProductController {
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := http.NewProductController(productService)
	return productController
}
