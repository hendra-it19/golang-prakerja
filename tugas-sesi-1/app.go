package main

import "fmt"

func main() {
	var name string = "Hendra"
	var age int = 23
	var status bool = true

	fmt.Printf("%s \n %d \n %t \n", name, age, status)
	fmt.Printf("%T \n %T \n %T \n", name, age, status)
}
