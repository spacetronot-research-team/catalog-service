package repository

import (
	"context"

	"github.com/spacetronot-research-team/catalog-service/internal/model"
	"gorm.io/gorm"
)

//go:generate mockgen -source=product.go -destination=mock/product.go -package=repository

type Product interface {
	// Create inserts product to db, return productID and error
	Create(ctx context.Context, product model.Product) (productID int64, err error)
	GetList()
	GetDetails()
	Update()
	// Delete will delete product from db by productID
	Delete(ctx context.Context, productID int64) (err error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) Product {
	return &productRepository{
		db: db,
	}
}

// Create inserts product to db, return productID and error
func (pr *productRepository) Create(ctx context.Context, product model.Product) (int64, error) {
	if err := pr.db.Create(&product).Error; err != nil {
		return 0, err
	}

	return product.ID, nil
}

// Delete will delete product from db by productID
func (pr *productRepository) Delete(ctx context.Context, productID int64) (err error) {
	query := pr.db.Debug().Where("id = ?", productID).Delete(&model.Product{})
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (*productRepository) GetDetails() {
	panic("unimplemented")
}

func (*productRepository) GetList() {
	panic("unimplemented")
}

func (*productRepository) Update() {
	panic("unimplemented")
}
