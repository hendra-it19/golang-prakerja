package controller

import (
	"net/http"
	"prakerja-tugas6/data/request"
	"prakerja-tugas6/data/response"
	"prakerja-tugas6/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(service service.ProductService) *ProductController {
	return &ProductController{productService: service}
}

func (controller *ProductController) FindAll(ctx *gin.Context) {
	data, err := controller.productService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Status: "success",
		Data:   data,
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller *ProductController) FindById(ctx *gin.Context) {
	taskId := ctx.Param("id")
	id, err := strconv.Atoi(taskId)

	data, err := controller.productService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Status: "success",
		Data:   data,
	}
	ctx.JSON(http.StatusOK, res)
}

func (controller *ProductController) Create(ctx *gin.Context) {
	req := request.CreateProductRequest{}
	ctx.ShouldBindJSON(&req)

	err := controller.productService.Create(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Status: "success",
		Data:   req,
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller *ProductController) Update(ctx *gin.Context) {
	req := request.UpdateProductRequest{}
	err := ctx.ShouldBindJSON(&req)

	taskId := ctx.Param("id")
	id, err := strconv.Atoi(taskId)

	_, err = controller.productService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	req.Id = uint(id)

	err = controller.productService.Update(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Status: "success",
		Data:   req,
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller *ProductController) Delete(ctx *gin.Context) {
	taskId := ctx.Param("id")
	id, err := strconv.Atoi(taskId)

	_, err = controller.productService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	err = controller.productService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Status: "success",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, res)
}
