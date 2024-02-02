package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spacetronot-research-team/catalog-service/pkg/middleware"
	"gorm.io/gorm"
)

func Add(ginEngine *gin.Engine, db *gorm.DB) {
	productController := getProductController(db)

	ginEngine.Use(middleware.Trace())

	ginEngine.POST("/api/products", productController.Create)
	ginEngine.DELETE("/api/products/:id", productController.Delete)
}
