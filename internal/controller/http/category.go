package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spacetronot-research-team/catalog-service/internal/dto"
	"github.com/spacetronot-research-team/catalog-service/internal/service"
	httpresponse "github.com/spacetronot-research-team/catalog-service/pkg/http_response"
)

type CategoryController struct {
	categoryService service.Category
}

func NewCategoryController(categoryService service.Category) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
	}
}

// Create inserts category to db, return categoryID and error
func (cc *CategoryController) Create(ctx *gin.Context) {
	dtoCategory := dto.CreateCategoryRequest{}

	if err := ctx.ShouldBindJSON(&dtoCategory); err != nil {
		logrus.WithContext(ctx).Error(err)
		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	categoryID, err := cc.categoryService.Create(ctx, dtoCategory)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"category_id": categoryID,
	}).Info("success inserts category to db")

	httpresponse.Write(ctx, http.StatusOK, categoryID, nil)
}
