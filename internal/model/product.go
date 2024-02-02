package model

type Product struct {
	ID         int64   `gorm:"column:id;primaryKey"`
	Name       string  `gorm:"column:name"`
	CategoryID int64   `gorm:"column:category_id"`
	Stock      int64   `gorm:"column:stock"`
	Price      float64 `gorm:"column:price"`
}
