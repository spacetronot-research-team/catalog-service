package service

import (
	"context"
	"testing"

	"github.com/spacetronot-research-team/catalog-service/internal/model"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"

	repository "github.com/spacetronot-research-team/catalog-service/internal/repository/mock"
)

func Test_productService_GetDetail(t *testing.T) {
	type fields struct {
		productRepository *repository.MockProduct
	}

	type args struct {
		ctx context.Context
		id  int
	}

	productSuccess := model.Product{
		ID:         1,
		CategoryId: 2,
		Category: model.Category{
			ID:   1,
			Name: "Handphone",
		},
		Name:  "iphone 14",
		Stock: 99,
		Price: 15000,
	}

	productFailed := model.Product{}

	tests := []struct {
		name        string
		mock        func(f fields)
		args        args
		wantProduct model.Product
		wantErr     bool
	}{
		{
			name: "Get Product Detail Success",
			mock: func(f fields) {
				f.productRepository.EXPECT().
					GetDetail(nil, 1).
					Return(productSuccess, nil)
			},
			args: args{
				ctx: nil,
				id:  1,
			},
			wantProduct: productSuccess,
			wantErr:     false,
		},
		{
			name: "Get Product Detail Failed",
			mock: func(f fields) {
				f.productRepository.EXPECT().
					GetDetail(nil, 1).
					Return(productFailed, nil)
			},
			args: args{
				ctx: nil,
				id:  1,
			},
			wantProduct: productFailed,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				productRepository: repository.NewMockProduct(ctrl),
			}
			tt.mock(f)

			ps := &productService{
				productRepository: f.productRepository,
			}
			gotProduct, err := ps.GetDetail(tt.args.ctx, tt.args.id)
			assert.Equal(t, tt.wantProduct, gotProduct)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
