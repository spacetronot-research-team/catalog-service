package dto

type CreateProductRequest struct {
	Name       string  `json:"name"`
	CategoryID int64   `json:"category_id"`
	Stock      int64   `json:"stock"`
	Price      float64 `json:"price"`
}

type GetProductListRequest struct {
	Limit  int64
	Offset int64
}

type Product struct {
	ID       int64
	Name     string
	Category Category
	Stock    int64
	Price    float64
}
