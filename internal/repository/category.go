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
	// GetList return categories
	GetList(ctx context.Context) (categories []model.Category, err error)
	GetDetails()
	// Update will update category by id for every field that is not default value
	Update(ctx context.Context, category model.Category) (categoryID int64, err error)
	// Delete will delete category from db by categoryID
	Delete(ctx context.Context, categoryID int64) (err error)
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

// Delete will delete category from db by categoryID
func (cr *categoryRepository) Delete(ctx context.Context, categoryID int64) error {
	query := cr.db.Where("id = ?", categoryID).Delete(&model.Category{})
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (*categoryRepository) GetDetails() {
	panic("unimplemented")
}

// GetList return categories
func (cr *categoryRepository) GetList(ctx context.Context) ([]model.Category, error) {
	categories := []model.Category{}
	if err := cr.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// Update will update category by id for every field that is not default value
func (cr *categoryRepository) Update(ctx context.Context, category model.Category) (int64, error) {
	query := cr.db.Updates(&category)
	if query.Error != nil {
		return 0, query.Error
	}
	if query.RowsAffected == 0 {
		return 0, gorm.ErrRecordNotFound
	}
	return category.ID, nil
}
