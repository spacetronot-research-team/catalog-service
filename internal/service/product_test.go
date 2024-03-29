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

func Test_productService_Delete(t *testing.T) {
	type fields struct {
		productRepository *repository.MockProduct
	}
	type args struct {
		ctx       context.Context
		productID int64
	}
	tests := []struct {
		name    string
		mock    func(f fields)
		args    args
		wantErr bool
	}{
		{
			name: "success delete product by id",
			mock: func(f fields) {
				f.productRepository.EXPECT().Delete(nil, int64(100)).Return(nil)
			},
			args: args{
				ctx:       nil,
				productID: 100,
			},
			wantErr: false,
		},
		{
			name: "fail delete product by id",
			mock: func(f fields) {
				f.productRepository.EXPECT().Delete(nil, int64(100)).Return(assert.AnError)
			},
			args: args{
				ctx:       nil,
				productID: 100,
			},
			wantErr: true,
		},
		{
			name: "product_id is not valid",
			mock: func(f fields) {},
			args: args{
				ctx:       nil,
				productID: -100,
			},
			wantErr: true,
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

			err := ps.Delete(tt.args.ctx, tt.args.productID)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func Test_productService_Update(t *testing.T) {
	type fields struct {
		productRepository *repository.MockProduct
	}
	type args struct {
		ctx        context.Context
		dtoProduct dto.UpdateProductRequest
	}
	tests := []struct {
		name    string
		mock    func(f fields)
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "validate error, all body is default value",
			mock: func(f fields) {},
			args: args{
				ctx: nil,
				dtoProduct: dto.UpdateProductRequest{
					ID:         0,
					Name:       "",
					CategoryID: 0,
					Stock:      0,
					Price:      0,
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "validate error, category_id is not valid",
			mock: func(f fields) {},
			args: args{
				ctx: nil,
				dtoProduct: dto.UpdateProductRequest{
					ID:         0,
					Name:       "",
					CategoryID: -11,
					Stock:      0,
					Price:      0,
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "validate error, stock is not valid",
			mock: func(f fields) {},
			args: args{
				ctx: nil,
				dtoProduct: dto.UpdateProductRequest{
					ID:         0,
					Name:       "",
					CategoryID: 0,
					Stock:      -11,
					Price:      0,
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "validate error, price is not valid",
			mock: func(f fields) {},
			args: args{
				ctx: nil,
				dtoProduct: dto.UpdateProductRequest{
					ID:         0,
					Name:       "",
					CategoryID: 0,
					Stock:      0,
					Price:      -11,
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "update from repo error",
			mock: func(f fields) {
				f.productRepository.EXPECT().Update(nil, model.Product{
					ID:         11,
					Name:       "new product name",
					CategoryID: 0,
					Stock:      0,
					Price:      0,
				}).Return(int64(0), assert.AnError)
			},
			args: args{
				ctx: nil,
				dtoProduct: dto.UpdateProductRequest{
					ID:         11,
					Name:       "new product name",
					CategoryID: 0,
					Stock:      0,
					Price:      0,
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "update product success",
			mock: func(f fields) {
				f.productRepository.EXPECT().Update(nil, model.Product{
					ID:         11,
					Name:       "new product name",
					CategoryID: 0,
					Stock:      0,
					Price:      0,
				}).Return(int64(89798), nil)
			},
			args: args{
				ctx: nil,
				dtoProduct: dto.UpdateProductRequest{
					ID:         11,
					Name:       "new product name",
					CategoryID: 0,
					Stock:      0,
					Price:      0,
				},
			},
			want:    89798,
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

			got, err := ps.Update(tt.args.ctx, tt.args.dtoProduct)
			assert.Equal(t, got, tt.want)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func Test_productService_GetList(t *testing.T) {
	type fields struct {
		productRepository *repository.MockProduct
	}
	type args struct {
		ctx   context.Context
		page  int
		limit int
	}
	tests := []struct {
		name    string
		mock    func(f fields)
		args    args
		want    []model.Product
		wantErr bool
	}{
		{
			name: "request invalid page less than 0",
			mock: func(f fields) {
			},
			args: args{
				ctx:   nil,
				page:  0,
				limit: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "request invalid limit less than 0",
			mock: func(f fields) {
			},
			args: args{
				ctx:   nil,
				page:  1,
				limit: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fail get products from repo",
			mock: func(f fields) {
				f.productRepository.EXPECT().
					GetList(nil, gomock.Any()).
					Return(nil, assert.AnError)
			},
			args: args{
				ctx:   nil,
				page:  10,
				limit: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success get products",
			mock: func(f fields) {
				f.productRepository.EXPECT().
					GetList(nil, gomock.Any()).
					Return([]model.Product{{}, {}}, nil)
			},
			args: args{
				ctx:   nil,
				page:  10,
				limit: 1,
			},
			want:    []model.Product{{}, {}},
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

			got, err := ps.GetList(tt.args.ctx, tt.args.page, tt.args.limit)
			assert.Equal(t, got, tt.want)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

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
