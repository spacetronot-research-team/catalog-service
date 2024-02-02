package service

import (
	"context"
	"testing"

	"github.com/spacetronot-research-team/catalog-service/internal/dto"
	"github.com/spacetronot-research-team/catalog-service/internal/model"
	repository "github.com/spacetronot-research-team/catalog-service/internal/repository/mock"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func Test_productService_Create(t *testing.T) {
	type fields struct {
		productRepository *repository.MockProduct
	}
	type args struct {
		ctx        context.Context
		dtoProduct dto.CreateProductRequest
	}
	tests := []struct {
		name    string
		mock    func(f fields)
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "category_id invalid",
			mock: func(f fields) {},
			args: args{
				ctx: nil,
				dtoProduct: dto.CreateProductRequest{
					Name:       "",
					CategoryID: -1,
					Stock:      100,
					Price:      100000,
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "stock invalid",
			mock: func(f fields) {},
			args: args{
				ctx: nil,
				dtoProduct: dto.CreateProductRequest{
					Name:       "",
					CategoryID: 1,
					Stock:      -100,
					Price:      100000,
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "price invalid",
			mock: func(f fields) {},
			args: args{
				ctx: nil,
				dtoProduct: dto.CreateProductRequest{
					Name:       "",
					CategoryID: 1,
					Stock:      100,
					Price:      -100000,
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "fail insert to db",
			mock: func(f fields) {
				f.productRepository.EXPECT().Create(nil, model.Product{
					Name:       "iphone 11",
					CategoryID: 1,
					Stock:      100,
					Price:      100000,
				}).Return(int64(0), assert.AnError)
			},
			args: args{
				ctx: nil,
				dtoProduct: dto.CreateProductRequest{
					Name:       "iphone 11",
					CategoryID: 1,
					Stock:      100,
					Price:      100000,
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "success insert to db",
			mock: func(f fields) {
				f.productRepository.EXPECT().Create(nil, model.Product{
					Name:       "iphone 11",
					CategoryID: 1,
					Stock:      100,
					Price:      100000,
				}).Return(int64(2), nil)
			},
			args: args{
				ctx: nil,
				dtoProduct: dto.CreateProductRequest{
					Name:       "iphone 11",
					CategoryID: 1,
					Stock:      100,
					Price:      100000,
				},
			},
			want:    2,
			wantErr: false,
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

			got, err := ps.Create(tt.args.ctx, tt.args.dtoProduct)
			assert.Equal(t, got, tt.want)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
