package category

import "errors"

var (
	ErrCategoryAlreadyExists = errors.New("category already exists")
)

type CreateCategoryRequest struct {
	Name string
}

type CreateCategoryResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
