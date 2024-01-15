package repository

import (
	"context"

	"github.com/spacetronot-research-team/catalog-service/internal/model"
	"gorm.io/gorm"
)

//go:generate mockgen -source=category.go -destination=mock/category.go -package=repository

type Category interface {
	// Create inserts category to db, return categoryID and error
	Create(ctx context.Context, category model.Category) (categoryID int64, err error)
	GetList()
	GetDetails()
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

// Create inserts category to db, return categoryID and error
func (cr *categoryRepository) Create(ctx context.Context, category model.Category) (int64, error) {
	if err := cr.db.Create(&category).Error; err != nil {
		return 0, err
	}

	return category.ID, nil
}

func (*categoryRepository) Delete() {
	panic("unimplemented")
}

func (*categoryRepository) GetDetails() {
	panic("unimplemented")
}

func (*categoryRepository) GetList() {
	panic("unimplemented")
}

func (*categoryRepository) Update() {
	panic("unimplemented")
}
