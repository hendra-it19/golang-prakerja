package repository

import (
	"errors"
	"prakerja-tugas6/data/request"
	"prakerja-tugas6/model"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductRepositoryImpl(Db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{Db: Db}
}

func (t ProductRepositoryImpl) FindAll() (product []model.Product, err error) {
	results := t.Db.Find(&product)
	if results.Error != nil {
		return nil, results.Error
	}

	return product, nil
}

func (t ProductRepositoryImpl) FindById(productID int) (product model.Product, err error) {
	result := t.Db.Find(&product, productID)

	if result.Error != nil {
		return model.Product{}, result.Error
	}

	if result.RowsAffected == 0 {
		return model.Product{}, errors.New("product is not found")
	}

	return product, nil
}

func (t *ProductRepositoryImpl) Save(product model.Product) (p model.Product, err error) {
	result := t.Db.Create(&product)
	if result.Error != nil {
		return p, result.Error
	}

	return p, nil
}

func (t *ProductRepositoryImpl) Update(product model.Product) (p model.Product, err error) {
	var data = request.UpdateProductRequest{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
	}

	result := t.Db.Model(&product).Updates(data)
	if result.Error != nil {
		return p, result.Error
	}

	return p, nil
}

func (t *ProductRepositoryImpl) Delete(productId int) error {
	var product model.Product

	result := t.Db.Where("id = ?", productId).Delete(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
