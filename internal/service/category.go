package service

import "github.com/spacetronot-research-team/catalog-service/internal/repository"

type Category interface {
	Create()
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

// Create implements Category.
func (*categoryService) Create() {
	panic("unimplemented")
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
