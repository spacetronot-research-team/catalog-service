package repository

type Category interface {
	Create()
	GetList()
	GetDetails()
	Update()
	Delete()
}

type categoryRepository struct {
}

func NewCategoryRepository() Category {
	return &categoryRepository{}
}

func (*categoryRepository) Create() {
	panic("unimplemented")
}

func (*categoryRepository) Delete() {
	panic("unimplemented")
}

func (*categoryRepository) GetDetails() {
	panic("unimplemented")
}

func (*categoryRepository) GetList() {
	panic("unimplemented")
}

func (*categoryRepository) Update() {
	panic("unimplemented")
}
