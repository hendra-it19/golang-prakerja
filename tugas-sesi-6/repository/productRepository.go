package repository

import (
	"prakerja-tugas6/model"
)

type ProductRepository interface {
	FindAll() ([]model.Product, error)
	FindById(productId int) (product model.Product, err error)
	Save(product model.Product) (p model.Product, err error)
	Update(product model.Product) (p model.Product, err error)
	Delete(productId int) error
}
