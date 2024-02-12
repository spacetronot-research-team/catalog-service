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

	ginEngine.POST("/api/products", productController.Create)
	ginEngine.GET("/api/products/:id", productController.GetDetail)
	ginEngine.DELETE("/api/products/:id", productController.Delete)
	ginEngine.PUT("/api/products/:id", productController.Update)
	ginEngine.GET("/api/products", productController.GetList)

	ginEngine.POST("/api/categories", categoryController.Create)
	ginEngine.GET("/api/categories/:id", categoryController.GetDetail)
	ginEngine.DELETE("/api/categories/:id", categoryController.Delete)
	ginEngine.PUT("/api/categories/:id", categoryController.Update)
	ginEngine.GET("/api/categories", categoryController.GetList)
}
