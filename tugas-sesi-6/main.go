package main

import (
	"log"
	"net/http"
	"prakerja-tugas6/config"
	"prakerja-tugas6/controller"
	"prakerja-tugas6/model"
	"prakerja-tugas6/repository"
	"prakerja-tugas6/router"
	"prakerja-tugas6/service"
	"prakerja-tugas6/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

func main() {
	//Database
	db := config.DatabaseConnection()
	db.AutoMigrate(&model.Product{})
	validate := validator.New()

	//Init Repository
	productRepository := repository.NewProductRepositoryImpl(db)

	//Init Service
	productService, err := service.NewProductServiceImpl(productRepository, validate)
	if err != nil {
		// Handle error appropriately, such as logging and exiting
		log.Fatalf("Error initializing task service: %v", err)
	}

	//Init controller
	productController := controller.NewProductController(productService)

	//Router
	routes := router.ProductRouter(productController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = server.ListenAndServe()
	utils.ErrorPanic(err)

}
