package dto

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type GetCategoryListRequest struct {
	Limit  int64
	Offset int64
}

type UpdateCategoryRequest struct {
	ID   int64  `json:"-"`
	Name string `json:"name"`
}

type Category struct {
	ID   int64
	Name string
}
