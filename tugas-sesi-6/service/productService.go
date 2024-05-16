package service

import (
	"prakerja-tugas6/data/request"
	"prakerja-tugas6/data/response"
)

type ProductService interface {
	Create(product request.CreateProductRequest) error
	Update(product request.UpdateProductRequest) error
	Delete(productId int) error
	FindById(productId int) (product response.ProductResponse, err error)
	FindAll() (product []response.ProductResponse, err error)
}
