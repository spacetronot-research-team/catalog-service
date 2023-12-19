package dto

import "github.com/spacetronot-research-team/catalog-service/internal/model"

type CreateCategoryRequest struct {
}

type CreateCategoryResponse struct {
	ID   string
	Name string
}

func (c model.Category) ToResponse() CreateCategoryResponse {
	return CreateCategoryResponse{
		ID:   c.ID,
		Name: c.Name,
	}
}
