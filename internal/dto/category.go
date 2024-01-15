package dto

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type GetCategoryListRequest struct {
	Limit  int64
	Offset int64
}

type UpdateCategoryRequest struct {
	Name string
}

type Category struct {
	ID   int64
	Name string
}
