package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

// Get detail category to db
func (pc *CategoryController) GetDetail(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	category, err := pc.categoryService.GetDetail(ctx, id)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"category": category,
	}).Info("success get detail category to db")

	httpresponse.Write(ctx, http.StatusOK, category, nil)
}
