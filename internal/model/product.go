package model

type Name string

const (
	NAMA_1 Name = "NAMA_1"
)

type Product struct {
	ID         string `gorm:"primary"`
	Name       Name
	CategoryID int
	Stock      int
	Price      Price
}
