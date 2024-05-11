package main

import "fmt"

// method dimiliki oleh suatu tipe data (func yang dimiliki oleh tipe data)
// func tidak memiliki tipe data sehingga bisa langsung digunakan

type Person struct {
	Name string
	Age  int
}
type Food struct {
	Name   string
	Price  int
	Origin string
}

// func
func Greet() {

}

// method
func (p Person) Greet() {
	fmt.Printf("%+v \n", p)
}
func (f Food) Introduce() {

}

func main() {
	var p1 Person = Person{
		Name: "Hendra",
		Age:  23,
	}
	p1.Greet()

	var p2 Person = Person{
		Name: "Lusi",
		Age:  21,
	}
	p2.Greet()

}
