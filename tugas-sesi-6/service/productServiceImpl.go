package service

import (
	"errors"
	"prakerja-tugas6/data/request"
	"prakerja-tugas6/data/response"
	"prakerja-tugas6/model"
	"prakerja-tugas6/repository"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	Validate          *validator.Validate
}

func NewProductServiceImpl(productRepository repository.ProductRepository, validate *validator.Validate) (service ProductService, err error) {
	if validate == nil {
		return nil, errors.New("validator instance cannot be nil")
	}
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		Validate:          validate,
	}, err
}

func (t ProductServiceImpl) FindAll() (products []response.ProductResponse, err error) {
	result, err := t.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}

	for _, value := range result {
		product := response.ProductResponse{
			Id:    value.Id,
			Name:  value.Name,
			Price: value.Price,
		}
		products = append(products, product)
	}
	return products, nil
}

func (t ProductServiceImpl) FindById(productId int) (response.ProductResponse, error) {
	data, err := t.ProductRepository.FindById(productId)
	if err != nil {
		return response.ProductResponse{}, err
	}

	res := response.ProductResponse{
		Id:    data.Id,
		Name:  data.Name,
		Price: data.Price,
	}
	return res, nil
}

func (t *ProductServiceImpl) Create(product request.CreateProductRequest) (err error) {
	err = t.Validate.Struct(product)

	if err != nil {
		return err
	}

	m := model.Product{
		Name:  product.Name,
		Price: product.Price,
	}

	t.ProductRepository.Save(m)

	return nil
}

func (t *ProductServiceImpl) Update(product request.UpdateProductRequest) (err error) {
	data, err := t.ProductRepository.FindById(int(product.Id))

	if err != nil {
		return err
	}

	data.Name = product.Name
	data.Price = product.Price
	t.ProductRepository.Update(data)
	return nil
}

func (t *ProductServiceImpl) Delete(productId int) (err error) {
	err = t.ProductRepository.Delete(productId)

	if err != nil {
		return err
	}

	return nil
}
