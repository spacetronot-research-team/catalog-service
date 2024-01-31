package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spacetronot-research-team/catalog-service/pkg/middleware"
	"gorm.io/gorm"
)

func Add(ginEngine *gin.Engine, db *gorm.DB) {
	productController := productController(db)
	categoryController := categoryController(db)
	ginEngine.Use(middleware.Trace())

	ginEngine.GET("/api/products/:id", productController.GetDetail)
	ginEngine.GET("/api/categories/:id", categoryController.GetDetail)
}
