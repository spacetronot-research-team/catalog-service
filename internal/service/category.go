package service

import (
	"context"

	"github.com/spacetronot-research-team/catalog-service/internal/model"
	"github.com/spacetronot-research-team/catalog-service/internal/repository"
)

type Category interface {
	Create()
	GetList()
	GetDetail(ctx context.Context, id int) (category model.Category, err error)
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

// Create implements Category.
func (*categoryService) Create() {
	panic("unimplemented")
}

// Delete implements Category.
func (*categoryService) Delete() {
	panic("unimplemented")
}

// GetDetails implements Category.
func (cs *categoryService) GetDetail(ctx context.Context, id int) (category model.Category, err error) {
	return cs.categoryRepository.GetDetail(ctx, id)
}

// GetList implements Category.
func (*categoryService) GetList() {
	panic("unimplemented")
}

// Update implements Category.
func (*categoryService) Update() {
	panic("unimplemented")
}
