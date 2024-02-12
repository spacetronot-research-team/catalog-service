package http

import (
	"fmt"
	"net/http"
	"strconv"

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

// Delete will delete category from db by categoryID
func (cc *CategoryController) Delete(ctx *gin.Context) {
	paramID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"id": ctx.Param("id"),
		}).Error(fmt.Errorf("fail convert param id string to int: %v", err))

		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	categoryID := int64(paramID)
	if err := cc.categoryService.Delete(ctx, categoryID); err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"category_id": categoryID,
		}).Error(err)

		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"category_id": categoryID,
	}).Info("success delete category from db")

	httpresponse.Write(ctx, http.StatusOK, categoryID, nil)
}

// GetList return categories
func (cc *CategoryController) GetList(ctx *gin.Context) {
	categories, err := cc.categoryService.GetList(ctx)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	logrus.WithContext(ctx).Info("success get categories")

	httpresponse.Write(ctx, http.StatusOK, categories, nil)
}

// Update will update category by id for every field that is not default value
func (cc *CategoryController) Update(ctx *gin.Context) {
	dtoCategory := dto.UpdateCategoryRequest{}
	if err := ctx.ShouldBindJSON(&dtoCategory); err != nil {
		logrus.WithContext(ctx).Error(err)
		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	paramID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"id": ctx.Param("id"),
		}).Error(fmt.Errorf("fail convert param id string to int: %v", err))

		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	dtoCategory.ID = int64(paramID)

	categoryID, err := cc.categoryService.Update(ctx, dtoCategory)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"category_id": categoryID,
	}).Info("success update category")

	httpresponse.Write(ctx, http.StatusOK, categoryID, nil)
}
