package service

import (
	"context"
	"errors"
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
	// Delete will delete category from db by categoryID
	Delete(ctx context.Context, categoryID int64) (err error)
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

// Delete will delete category from db by categoryID
func (cs *categoryService) Delete(ctx context.Context, categoryID int64) error {
	if err := cs.validateDeleteCategoryID(categoryID); err != nil {
		return fmt.Errorf("request invalid: %v", err)
	}

	if err := cs.categoryRepository.Delete(ctx, categoryID); err != nil {
		return fmt.Errorf("fail delete category by id: %v", err)
	}

	return nil
}

func (*categoryService) validateDeleteCategoryID(categoryID int64) error {
	if categoryID <= 0 {
		return errors.New("category_id is not valid")
	}
	return nil
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
