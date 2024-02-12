package service

import (
	"context"
	"testing"

	"github.com/spacetronot-research-team/catalog-service/internal/dto"
	"github.com/spacetronot-research-team/catalog-service/internal/model"
	repository "github.com/spacetronot-research-team/catalog-service/internal/repository/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_categoryService_Create(t *testing.T) {
	type fields struct {
		categoryRepository *repository.MockCategory
	}
	type args struct {
		ctx         context.Context
		dtoCategory dto.CreateCategoryRequest
	}
	tests := []struct {
		name    string
		mock    func(f fields)
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "fail insert category",
			mock: func(f fields) {
				f.categoryRepository.EXPECT().
					Create(nil, model.Category{Name: "foo"}).
					Return(int64(0), assert.AnError)
			},
			args: args{
				ctx: nil,
				dtoCategory: dto.CreateCategoryRequest{
					Name: "foo",
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "success insert category",
			mock: func(f fields) {
				f.categoryRepository.EXPECT().
					Create(nil, model.Category{Name: "foo"}).
					Return(int64(100), nil)
			},
			args: args{
				ctx: nil,
				dtoCategory: dto.CreateCategoryRequest{
					Name: "foo",
				},
			},
			want:    100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				categoryRepository: repository.NewMockCategory(ctrl),
			}
			tt.mock(f)

			cs := &categoryService{
				categoryRepository: f.categoryRepository,
			}

			got, err := cs.Create(tt.args.ctx, tt.args.dtoCategory)
			assert.Equal(t, got, tt.want)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

		})
	}
}

func Test_categoryService_Delete(t *testing.T) {
	type fields struct {
		categoryRepository *repository.MockCategory
	}
	type args struct {
		ctx        context.Context
		categoryID int64
	}
	tests := []struct {
		name    string
		mock    func(f fields)
		args    args
		wantErr bool
	}{
		{
			name: "success delete category by id",
			mock: func(f fields) {
				f.categoryRepository.EXPECT().Delete(nil, int64(100)).Return(nil)
			},
			args: args{
				ctx:        nil,
				categoryID: 100,
			},
			wantErr: false,
		},
		{
			name: "fail delete category by id",
			mock: func(f fields) {
				f.categoryRepository.EXPECT().Delete(nil, int64(100)).Return(assert.AnError)
			},
			args: args{
				ctx:        nil,
				categoryID: 100,
			},
			wantErr: true,
		},
		{
			name: "category_id is not valid",
			mock: func(f fields) {},
			args: args{
				ctx:        nil,
				categoryID: -100,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				categoryRepository: repository.NewMockCategory(ctrl),
			}
			tt.mock(f)

			cs := &categoryService{
				categoryRepository: f.categoryRepository,
			}
			err := cs.Delete(tt.args.ctx, tt.args.categoryID)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func Test_categoryService_GetList(t *testing.T) {
	type fields struct {
		categoryRepository *repository.MockCategory
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		mock    func(f fields)
		args    args
		want    []model.Category
		wantErr bool
	}{
		{
			name: "fail get category list from repo",
			mock: func(f fields) {
				f.categoryRepository.EXPECT().
					GetList(nil).
					Return(nil, assert.AnError)
			},
			args: args{
				ctx: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success get category list",
			mock: func(f fields) {
				f.categoryRepository.EXPECT().
					GetList(nil).
					Return([]model.Category{{}, {}, {}}, nil)
			},
			args: args{
				ctx: nil,
			},
			want:    []model.Category{{}, {}, {}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				categoryRepository: repository.NewMockCategory(ctrl),
			}
			tt.mock(f)

			cs := &categoryService{
				categoryRepository: f.categoryRepository,
			}
			got, err := cs.GetList(tt.args.ctx)
			assert.Equal(t, got, tt.want)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func Test_categoryService_Update(t *testing.T) {
	type fields struct {
		categoryRepository *repository.MockCategory
	}
	type args struct {
		ctx         context.Context
		dtoCategory dto.UpdateCategoryRequest
	}
	tests := []struct {
		name    string
		mock    func(f fields)
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "request invalid, all body is default value",
			mock: func(f fields) {},
			args: args{
				ctx: nil,
				dtoCategory: dto.UpdateCategoryRequest{
					ID:   0,
					Name: "",
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "fail update from repo",
			mock: func(f fields) {
				f.categoryRepository.EXPECT().
					Update(nil, gomock.Any()).
					Return(int64(0), assert.AnError)
			},
			args: args{
				ctx: nil,
				dtoCategory: dto.UpdateCategoryRequest{
					ID:   2,
					Name: "categoryfoo",
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "success update category",
			mock: func(f fields) {
				f.categoryRepository.EXPECT().
					Update(nil, gomock.Any()).
					Return(int64(23342), nil)
			},
			args: args{
				ctx: nil,
				dtoCategory: dto.UpdateCategoryRequest{
					ID:   2,
					Name: "categoryfoo",
				},
			},
			want:    23342,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				categoryRepository: repository.NewMockCategory(ctrl),
			}
			tt.mock(f)

			cs := &categoryService{
				categoryRepository: f.categoryRepository,
			}

			got, err := cs.Update(tt.args.ctx, tt.args.dtoCategory)
			assert.Equal(t, got, tt.want)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func Test_categoryService_GetDetail(t *testing.T) {
	type fields struct {
		categoryRepository *repository.MockCategory
	}
	type args struct {
		ctx context.Context
		id  int
	}

	categorySuccess := model.Category{
		ID:   1,
		Name: "Handphone",
	}

	categoryFailed := model.Category{}

	tests := []struct {
		name         string
		mock         func(f fields)
		args         args
		wantCategory model.Category
		wantErr      bool
	}{
		{
			name: "Get Category Detail Success",
			mock: func(f fields) {
				f.categoryRepository.EXPECT().
					GetDetail(nil, 1).
					Return(categorySuccess, nil)
			},
			args: args{
				ctx: nil,
				id:  1,
			},
			wantCategory: categorySuccess,
			wantErr:      false,
		},
		{
			name: "Get Product Detail Failed",
			mock: func(f fields) {
				f.categoryRepository.EXPECT().
					GetDetail(nil, 1).
					Return(categoryFailed, nil)
			},
			args: args{
				ctx: nil,
				id:  1,
			},
			wantCategory: categoryFailed,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				categoryRepository: repository.NewMockCategory(ctrl),
			}
			tt.mock(f)

			cs := &categoryService{
				categoryRepository: f.categoryRepository,
			}
			gotCategory, err := cs.GetDetail(tt.args.ctx, tt.args.id)
			assert.Equal(t, tt.wantCategory, gotCategory)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
