package repository

import (
	"context"

	"github.com/spacetronot-research-team/catalog-service/internal/model"
	"gorm.io/gorm"
)

type Category interface {
	Create()
	GetList()
	GetDetail(ctx context.Context, id int) (category model.Category, err error)
	Update()
	Delete()
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) Category {
	return &categoryRepository{
		db: db,
	}
}

func (*categoryRepository) Create() {
	panic("unimplemented")
}

func (*categoryRepository) Delete() {
	panic("unimplemented")
}

func (cr *categoryRepository) GetDetail(ctx context.Context, id int) (category model.Category, err error) {
	query := cr.db.
		First(&category, id)
	if query.Error != nil {
		return model.Category{}, query.Error
	}
	return
}

func (*categoryRepository) GetList() {
	panic("unimplemented")
}

func (*categoryRepository) Update() {
	panic("unimplemented")
}
