package repository

type Product interface {
	Create()
	GetList()
	GetDetails()
	Update()
	Delete()
}

type productRepository struct {
}

func NewProductRepository() Product {
	return &productRepository{}
}

func (*productRepository) Create() {
	panic("unimplemented")
}

func (*productRepository) Delete() {
	panic("unimplemented")
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
