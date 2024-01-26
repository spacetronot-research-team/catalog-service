package repository

import (
	"context"

	"github.com/spacetronot-research-team/catalog-service/internal/model"
	"gorm.io/gorm"
)

type Product interface {
	Create()
	GetList()
	GetDetail(ctx context.Context, id int) (product model.Product, err error)
	Update()
	Delete()
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) Product {
	return &productRepository{
		db: db,
	}
}

func (*productRepository) Create() {
	panic("unimplemented")
}

func (*productRepository) Delete() {
	panic("unimplemented")
}

func (pr *productRepository) GetDetail(ctx context.Context, id int) (product model.Product, err error) {
	query := pr.db.
		Joins("Category").
		First(&product, id)
	if query.Error != nil {
		return model.Product{}, query.Error
	}
	return
}

func (*productRepository) GetList() {
	panic("unimplemented")
}

func (*productRepository) Update() {
	panic("unimplemented")
}
