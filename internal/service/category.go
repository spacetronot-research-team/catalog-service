package service

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/spacetronot-research-team/catalog-service/internal/dto"
	"github.com/spacetronot-research-team/catalog-service/internal/model"
	"github.com/spacetronot-research-team/catalog-service/internal/repository"
)

type Category interface {
	// Create inserts category to db, return categoryID and error
	Create(ctx context.Context, dtoCategory dto.CreateCategoryRequest) (categoryID int64, err error)
	// GetList return categories
	GetList(ctx context.Context) (categories []model.Category, err error)
	GetDetails()
	// Update will update category by id for every field that is not default value
	Update(ctx context.Context, dtoCategory dto.UpdateCategoryRequest) (categoryID int64, err error)
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

// GetList return categories
func (cs *categoryService) GetList(ctx context.Context) ([]model.Category, error) {
	categories, err := cs.categoryRepository.GetList(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail get category list: %v", err)
	}
	return categories, nil
}

// Update will update category by id for every field that is not default value
func (cs *categoryService) Update(ctx context.Context, dtoCategory dto.UpdateCategoryRequest) (int64, error) {
	if err := cs.validateDTOUpdateCategoryRequest(dtoCategory); err != nil {
		return 0, fmt.Errorf("update category request invalid: %v", err)
	}

	category := cs.dtoUpdateCategoryRequestToModelCategory(dtoCategory)

	categoryID, err := cs.categoryRepository.Update(ctx, category)
	if err != nil {
		return 0, fmt.Errorf("fail update category: %v", err)
	}

	return categoryID, nil
}

func (cs *categoryService) validateDTOUpdateCategoryRequest(dtoCategory dto.UpdateCategoryRequest) error {
	dtoCategoryDeultValueBody := dto.UpdateCategoryRequest{ID: dtoCategory.ID} // ID got from url param, not json body
	isAllBodyDefaultValue := reflect.DeepEqual(&dtoCategory, &dtoCategoryDeultValueBody)
	if isAllBodyDefaultValue {
		return errors.New("nothing to be update")
	}
	return nil
}

func (cs *categoryService) dtoUpdateCategoryRequestToModelCategory(dtoCategory dto.UpdateCategoryRequest) model.Category {
	return model.Category{
		ID:   dtoCategory.ID,
		Name: dtoCategory.Name,
	}
}
