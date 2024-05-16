package router

import (
	"prakerja-tugas6/controller"

	"github.com/gin-gonic/gin"
)

func ProductRouter(productController *controller.ProductController) *gin.Engine {
	service := gin.Default()

	router := service.Group("/products")

	router.GET("", productController.FindAll)
	router.GET("/:id", productController.FindById)
	router.POST("", productController.Create)
	router.PUT("/:id", productController.Update)
	router.DELETE("/:id", productController.Delete)

	return service
}
