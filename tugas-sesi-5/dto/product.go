package dto

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type GetProductResponse struct {
	Status string    `json:"status"`
	Data   []Product `json:"data"`
}

type CreateProductRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type CreateProductResponse struct {
	Status string  `json:"status"`
	Data   Product `json:"data"`
}

type UpdateProductRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type UpdateProductResponse struct {
	Status string  `json:"status"`
	Data   Product `json:"data"`
}

type DeleteProductResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}
