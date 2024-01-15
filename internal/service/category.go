package service

import (
	"context"
	"fmt"

	"github.com/spacetronot-research-team/catalog-service/internal/dto"
	"github.com/spacetronot-research-team/catalog-service/internal/model"
	"github.com/spacetronot-research-team/catalog-service/internal/repository"
)

type Category interface {
	// Create inserts category to db, return categoryID and error
	Create(ctx context.Context, dtoCategory dto.CreateCategoryRequest) (categoryID int64, err error)
	GetList()
	GetDetails()
	Update()
	Delete()
}

type categoryService struct {
	categoryRepository repository.Category
}

func NewCategoryService(categoryRepository repository.Category) Category {
	return &categoryService{
		categoryRepository: categoryRepository,
	}
}

// Create inserts category to db, return categoryID and error
func (cs *categoryService) Create(ctx context.Context, dtoCategory dto.CreateCategoryRequest) (int64, error) {
	categoryID, err := cs.categoryRepository.Create(ctx, model.Category{Name: dtoCategory.Name})
	if err != nil {
		return 0, fmt.Errorf("fail insert category to db: %v", err)
	}
	return categoryID, nil
}

// Delete implements Category.
func (*categoryService) Delete() {
	panic("unimplemented")
}

// GetDetails implements Category.
func (*categoryService) GetDetails() {
	panic("unimplemented")
}

// GetList implements Category.
func (*categoryService) GetList() {
	panic("unimplemented")
}

// Update implements Category.
func (*categoryService) Update() {
	panic("unimplemented")
}
