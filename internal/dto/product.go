package dto

type CreateProductRequest struct {
	Name       string
	CategoryID int64
	Stock      int64
	Price      float64
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
