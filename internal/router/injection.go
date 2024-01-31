package router

import (
	"github.com/spacetronot-research-team/catalog-service/internal/controller/http"
	"github.com/spacetronot-research-team/catalog-service/internal/repository"
	"github.com/spacetronot-research-team/catalog-service/internal/service"
	"gorm.io/gorm"
)

func productController(db *gorm.DB) *http.ProductController {
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := http.NewProductController(productService)
	return productController
}

func categoryController(db *gorm.DB) *http.CategoryController {
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := http.NewCategoryController(categoryService)
	return categoryController
}
