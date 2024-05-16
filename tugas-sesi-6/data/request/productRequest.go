package request

type CreateProductRequest struct {
	Name  string `validate:"required" json:"name"`
	Price uint   `validate:"required" json:"price"`
}

type UpdateProductRequest struct {
	Id    uint   `validate:"required"`
	Name  string `validate:"required,max=200" json:"name"`
	Price uint   `validate:"required" json:"price"`
}
