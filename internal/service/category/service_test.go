package category

import (
	"context"
	"testing"

	"github.com/spacetronot-research-team/catalog-service/internal/entity"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	persistenceRepo := NewMockpersistenceRepository(ctrl)
	cacheRepo := NewMockcacheRepository(ctrl)

	type args struct {
		persistenceRepo persistenceRepository
		cacheRepo       cacheRepository
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{
			name: "Creates a new Service instance",
			args: args{
				persistenceRepo: persistenceRepo,
				cacheRepo:       cacheRepo,
			},
			want: &categoryService{
				persistenceRepo: persistenceRepo,
				cacheRepo:       cacheRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewService(tt.args.persistenceRepo, tt.args.cacheRepo)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_categoryService_CreateCategory(t *testing.T) {
	type fields struct {
		persistenceRepo *MockpersistenceRepository
		cacheRepo       *MockcacheRepository
	}
	type args struct {
		ctx context.Context
		req CreateCategoryRequest
	}
	tests := []struct {
		name    string
		args    args
		mock    func(f fields)
		want    CreateCategoryResponse
		wantErr error
	}{
		{
			name: "Success to creates a new category",
			args: args{
				ctx: context.Background(),
				req: CreateCategoryRequest{
					Name: "Laptop",
				},
			},
			mock: func(f fields) {
				f.persistenceRepo.EXPECT().FindCategoryByName(gomock.Any(), "Laptop").
					Return(entity.Category{}, entity.ErrCategoryNotFound)

				f.persistenceRepo.EXPECT().InsertCategory(gomock.Any(), entity.Category{Name: "Laptop"}).
					Return(int64(1), nil)
			},
			want: CreateCategoryResponse{
				ID:   1,
				Name: "Laptop",
			},
			wantErr: nil,
		},
		{
			name: "Failed to creates a new category, category already exists",
			args: args{
				ctx: context.Background(),
				req: CreateCategoryRequest{
					Name: "Smartphone",
				},
			},
			mock: func(f fields) {
				f.persistenceRepo.EXPECT().FindCategoryByName(gomock.Any(), "Smartphone").
					Return(entity.Category{
						ID:   int64(2),
						Name: "Smartphone",
					}, nil)
			},
			want:    CreateCategoryResponse{},
			wantErr: ErrCategoryAlreadyExists,
		},
		{
			name: "An error is raised when finding a category by name",
			args: args{
				ctx: context.Background(),
				req: CreateCategoryRequest{
					Name: "Smartphone",
				},
			},
			mock: func(f fields) {
				f.persistenceRepo.EXPECT().FindCategoryByName(gomock.Any(), "Smartphone").
					Return(entity.Category{}, assert.AnError)
			},
			want:    CreateCategoryResponse{},
			wantErr: assert.AnError,
		},
		{
			name: "An error is raised when inserting a new category",
			args: args{
				ctx: context.Background(),
				req: CreateCategoryRequest{
					Name: "Smartphone",
				},
			},
			mock: func(f fields) {
				f.persistenceRepo.EXPECT().FindCategoryByName(gomock.Any(), "Smartphone").
					Return(entity.Category{}, entity.ErrCategoryNotFound)

				f.persistenceRepo.EXPECT().InsertCategory(gomock.Any(), entity.Category{Name: "Smartphone"}).
					Return(int64(0), assert.AnError)
			},
			want:    CreateCategoryResponse{},
			wantErr: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			persistenceRepo := NewMockpersistenceRepository(ctrl)
			cacheRepo := NewMockcacheRepository(ctrl)

			f := fields{
				persistenceRepo: persistenceRepo,
				cacheRepo:       cacheRepo,
			}
			tt.mock(f)

			svc := &categoryService{
				persistenceRepo: f.persistenceRepo,
				cacheRepo:       f.cacheRepo,
			}
			got, err := svc.CreateCategory(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
