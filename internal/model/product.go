package model

type Product struct {
	ID         int64    `json:"id" gorm:"column:id;primaryKey"`
	CategoryId int64    `json:"category_id" gorm:"column:category_id"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryId"`
	Name       string   `json:"name" gorm:"column:name"`
	Stock      int64    `json:"stock" gorm:"column:stock"`
	Price      float64  `json:"price" gorm:"column:price"`
}
