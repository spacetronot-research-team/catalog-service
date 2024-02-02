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
