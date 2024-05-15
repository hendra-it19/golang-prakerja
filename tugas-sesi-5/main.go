package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"tugas-sesi-5/dto"
)

var products []dto.Product = []dto.Product{}

func GenerateProductId() int {
	if len(products) > 0 {
		productId := products[len(products)-1].Id
		id, _ := strconv.Atoi(productId)
		return id + 1
	}
	return 1
}

func GetProduct(w http.ResponseWriter) {
	response := dto.GetProductResponse{
		Status: "success",
		Data:   products,
	}
	b, err := json.Marshal(response)
	_ = err
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", 422)
		return
	}
	var requestBody dto.CreateProductRequest
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "Something went wrong!", 500)
		return
	}
	product := dto.Product{
		Id:    fmt.Sprintf("%d", GenerateProductId()),
		Name:  requestBody.Name,
		Price: requestBody.Price,
	}
	products = append(products, product)
	response := dto.CreateProductResponse{
		Status: "success",
		Data:   product,
	}
	b, err := json.Marshal(response)
	_ = err
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	index := 0
	for i, v := range products {
		if v.Id == productId {
			index = i
			break
		}
	}
	if index == 0 {
		http.Error(w, "product not found!", 404)
		return
	}
	products = append(products[:index], products[index+1:]...)
	response := dto.DeleteProductResponse{
		Status: "success",
		Data:   nil,
	}
	b, err := json.Marshal(response)
	_ = err
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// cari
	productId := r.URL.Query().Get("id")
	index := 0
	for i, v := range products {
		if v.Id == productId {
			index = i
			break
		}
	}
	if index == 0 {
		http.Error(w, "product not found!", 404)
		return
	}
	// ambil inputan
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", 422)
		return
	}
	var requestBody dto.UpdateProductRequest
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "Something went wrong!", 500)
		return
	}
	// update
	products[index].Name = requestBody.Name
	products[index].Price = requestBody.Price
	response := dto.UpdateProductResponse{
		Status: "success",
		Data:   products[index],
	}
	b, err := json.Marshal(response)
	_ = err
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)

}

func main() {

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			GetProduct(w)
			return
		}
		if r.Method == "POST" {
			CreateProduct(w, r)
			return
		}
		if r.Method == "DELETE" {
			DeleteProduct(w, r)
			return
		}
		if r.Method == "PUT" {
			UpdateProduct(w, r)
			return
		}
	})

	fmt.Println("Listening in PORT 8080")
	http.ListenAndServe(":8080", nil)
}
