package service

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/spacetronot-research-team/catalog-service/internal/dto"
	"github.com/spacetronot-research-team/catalog-service/internal/model"
	"github.com/spacetronot-research-team/catalog-service/internal/repository"
	"github.com/spacetronot-research-team/catalog-service/pkg/pagination"
)

type Product interface {
	// Create inserts product to db, return productID and error
	Create(ctx context.Context, dtoProduct dto.CreateProductRequest) (productID int64, err error)
	// GetList return products filtered using pagination
	GetList(ctx context.Context, page int, limit int) (products []model.Product, err error)
	GetDetails()
	// Update will update product by id for every field that is not default value
	Update(ctx context.Context, dtoProduct dto.UpdateProductRequest) (productID int64, err error)
	// Delete will delete product from db by productID
	Delete(ctx context.Context, productID int64) (err error)
}

type productService struct {
	productRepository repository.Product
}

func NewProductService(productRepository repository.Product) Product {
	return &productService{
		productRepository: productRepository,
	}
}

// Create inserts product to db, return productID and error
func (ps *productService) Create(ctx context.Context, dtoProduct dto.CreateProductRequest) (int64, error) {
	if err := ps.validateDTOCreateProductRequest(dtoProduct); err != nil {
		return 0, fmt.Errorf("create product request invalid: %v", err)
	}

	product := ps.dtoCreateProductRequestToModelProduct(dtoProduct)

	productID, err := ps.productRepository.Create(ctx, product)
	if err != nil {
		return 0, fmt.Errorf("fail inserts product to db: %v", err)
	}

	return productID, nil
}

func (ps *productService) validateDTOCreateProductRequest(dtoProduct dto.CreateProductRequest) error {
	if dtoProduct.CategoryID < 0 {
		return errors.New("category_id not found")
	}

	if dtoProduct.Stock < 0 {
		return errors.New("stock can not less than 0")
	}

	if dtoProduct.Price < 0 {
		return errors.New("price can not less than 0")
	}

	return nil
}

func (ps *productService) dtoCreateProductRequestToModelProduct(dtoProduct dto.CreateProductRequest) model.Product {
	return model.Product{
		Name:       dtoProduct.Name,
		CategoryID: dtoProduct.CategoryID,
		Stock:      dtoProduct.Stock,
		Price:      dtoProduct.Price,
	}
}

// Delete will delete product from db by productID
func (ps *productService) Delete(ctx context.Context, productID int64) (err error) {
	if err := ps.validateDeleteProductID(productID); err != nil {
		return fmt.Errorf("request invalid: %v", err)
	}

	if err := ps.productRepository.Delete(ctx, productID); err != nil {
		return fmt.Errorf("fail delete product by id: %v", err)
	}
	return nil
}

func (ps *productService) validateDeleteProductID(productID int64) error {
	if productID <= 0 {
		return errors.New("product_id is not valid")
	}
	return nil
}

// GetDetails implements Product.
func (*productService) GetDetails() {
	panic("unimplemented")
}

// GetList return products filtered using pagination
func (ps *productService) GetList(ctx context.Context, page int, limit int) ([]model.Product, error) {
	if err := ps.validateGetListArgs(page, limit); err != nil {
		return nil, fmt.Errorf("invalid request: %v", err)
	}

	pagination := pagination.New(page, limit)

	products, err := ps.productRepository.GetList(ctx, pagination)
	if err != nil {
		return nil, fmt.Errorf("fail get product list: %v", err)
	}

	return products, nil
}

func (ps *productService) validateGetListArgs(page int, limit int) error {
	if page < 1 {
		return errors.New("page can not less than 1")
	}
	if limit < 1 {
		return errors.New("limit can not less than 1")
	}
	return nil
}

// Update will update product by id for every field that is not default value
func (ps *productService) Update(ctx context.Context, dtoProduct dto.UpdateProductRequest) (int64, error) {
	if err := ps.validateDTOUpdateProductRequest(dtoProduct); err != nil {
		return 0, fmt.Errorf("update product request invalid: %v", err)
	}

	product := ps.dtoUpdateProductRequestToModelProduct(dtoProduct)

	productID, err := ps.productRepository.Update(ctx, product)
	if err != nil {
		return 0, fmt.Errorf("fail update product: %v", err)
	}

	return productID, nil
}

func (ps *productService) validateDTOUpdateProductRequest(dtoProduct dto.UpdateProductRequest) error {
	dtoProductDefaultValueBody := dto.UpdateProductRequest{ID: dtoProduct.ID} // ID got from url param, not json body
	isAllBodyDefaultValue := reflect.DeepEqual(&dtoProduct, &dtoProductDefaultValueBody)
	if isAllBodyDefaultValue {
		return errors.New("nothing to be update")
	}

	if dtoProduct.CategoryID < 0 {
		return errors.New("category_id not found")
	}

	if dtoProduct.Stock < 0 {
		return errors.New("stock can not less than 0")
	}

	if dtoProduct.Price < 0 {
		return errors.New("price can not less than 0")
	}

	return nil
}

func (ps *productService) dtoUpdateProductRequestToModelProduct(dtoProduct dto.UpdateProductRequest) model.Product {
	return model.Product{
		ID:         dtoProduct.ID,
		Name:       dtoProduct.Name,
		CategoryID: dtoProduct.CategoryID,
		Stock:      dtoProduct.Stock,
		Price:      dtoProduct.Price,
	}
}
