package entity

import "errors"

var (
	ErrCategoryNotFound      = errors.New("category does not exist")
	ErrCategoryAlreadyExists = errors.New("category already exists")
)

type Category struct {
	ID   int64
	Name string
}
