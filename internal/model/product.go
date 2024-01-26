package model

type Product struct {
	ID         int64    `json:"id"`
	CategoryId int64    `json:"category_id"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryId"`
	Name       string   `json:"name"`
	Stock      int64    `json:"stock"`
	Price      float64  `json:"price"`
}
