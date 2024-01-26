package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spacetronot-research-team/catalog-service/internal/service"
	httpresponse "github.com/spacetronot-research-team/catalog-service/pkg/http_response"
)

type ProductController struct {
	productService service.Product
}

func NewProductController(productService service.Product) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

// Get detail product to db
func (pc *ProductController) GetDetail(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	product, err := pc.productService.GetDetail(ctx, id)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"product": product,
	}).Info("success get detail product to db")

	httpresponse.Write(ctx, http.StatusOK, product, nil)
}
