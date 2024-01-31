package service

import (
	"context"
	"testing"

	"github.com/spacetronot-research-team/catalog-service/internal/model"
	repository "github.com/spacetronot-research-team/catalog-service/internal/repository/mock"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

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
