package dto

import "github.com/spacetronot-research-team/catalog-service/internal/model"

type CreateProductRequest struct {
	Amount int
}

type CreateProductResponse struct {
	Success bool
}

func (p model.Product) ToResponse() CreateProductResponse {
	return CreateProductResponse{
		ID:   p.ID,
		Name: p.Name,
	}
}
