package response

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ProductResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}
