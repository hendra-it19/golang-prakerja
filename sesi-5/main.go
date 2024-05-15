package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var student Student = Student{
		Id:   1,
		Name: "Hendra",
		Age:  23,
	}

	jsonByte, err := json.Marshal(student)
	if err != nil {
		panic(err)
	}
	_ = jsonByte
	// fmt.Println(jsonByte)
	// fmt.Println(string(jsonByte))

	var customJson string = `
	{"id":1,"name":"Hendra","age":20}
	`
	var student2 Student
	var sb []byte = []byte(customJson)
	err = json.Unmarshal(sb, &student2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v \n", student2)
}
