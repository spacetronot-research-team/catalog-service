package http

import (
	"github.com/spacetronot-research-team/catalog-service/internal/service"
)

type Handler struct {
	productService service.Product
}

func NewHTTPHandler(productService service.Product) *Handler {
	return &Handler{
		productService: productService,
	}
}

func (h *Handler) handleGetProductByID(id int64) {
	productResponse := h.productService.GetDetails()

	return ProtoResponse{
		ID:   productResponse.ID,
		Name: productResponse.Name,
	}
}
