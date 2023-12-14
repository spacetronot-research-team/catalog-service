package category

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spacetronot-research-team/catalog-service/internal/entity"
)

//go:generate mockgen -source=service.go -destination=service_mock.go -package=category

type persistenceRepository interface {
	InsertCategory(ctx context.Context, category entity.Category) (int64, error)
	FindCategoryByID(ctx context.Context, id int64) (entity.Category, error)
	FindCategoryByName(ctx context.Context, name string) (entity.Category, error)
	UpdateCategory(ctx context.Context, category entity.Category) error
	DeleteCategory(ctx context.Context, id int64) error
}

type cacheRepository interface {
	InsertCategory(ctx context.Context, category entity.Category, ttl time.Duration) error
}

type Service interface {
	CreateCategory(ctx context.Context, req CreateCategoryRequest) (CreateCategoryResponse, error)
}

type categoryService struct {
	persistenceRepo persistenceRepository
	cacheRepo       cacheRepository
}

func NewService(persistenceRepo persistenceRepository, cacheRepo cacheRepository) Service {
	return &categoryService{
		persistenceRepo: persistenceRepo,
		cacheRepo:       cacheRepo,
	}
}

func (svc *categoryService) CreateCategory(ctx context.Context, req CreateCategoryRequest) (CreateCategoryResponse, error) {
	_, err := svc.persistenceRepo.FindCategoryByName(ctx, req.Name)
	if err != nil && err != entity.ErrCategoryNotFound {
		log.WithFields(log.Fields{
			"name": req.Name,
		}).Errorf("svc.persistenceRepo.FindCategoryByName() got error: %s", err.Error())

		return CreateCategoryResponse{}, err
	}

	if err == nil {
		return CreateCategoryResponse{}, ErrCategoryAlreadyExists
	}

	categoryID, err := svc.persistenceRepo.InsertCategory(ctx, entity.Category{
		Name: req.Name,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"name": req.Name,
		}).Errorf("svc.persistenceRepo.InsertCategory() got error: %s", err.Error())

		return CreateCategoryResponse{}, err
	}

	response := CreateCategoryResponse{
		ID:   categoryID,
		Name: req.Name,
	}

	err = svc.cacheRepo.InsertCategory(ctx, entity.Category{
		ID:   response.ID,
		Name: response.Name,
	}, 2*time.Hour)
	if err != nil {
		log.WithFields(log.Fields{
			"id":   response.ID,
			"name": response.Name,
		}).Errorf("svc.cacheRepo.InsertCategory() got error: %s", err.Error())
	}

	return response, nil
}
