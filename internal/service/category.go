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
	return &categoryService{categoryRepository: categoryRepository}
}

func (*categoryService) Create() {
	panic("unimplemented")
}

func (*categoryService) Delete() {
	panic("unimplemented")
}

func (*categoryService) GetDetails() {
	panic("unimplemented")
}

func (*categoryService) GetList() {
	panic("unimplemented")
}

func (*categoryService) Update() {
	panic("unimplemented")
}
