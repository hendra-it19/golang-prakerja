package main

import "fmt"

func main() {
	// foods := []string{"Nasi Goreng", "Mie Ayam", "Ayam Goreng"}

	// newFoods := foods

	// fmt.Printf("food len(%d), cap(%d) \n", len(foods), cap(foods))
	// fmt.Printf("newFoods len(%d), cap(%d) \n", len(newFoods), cap(newFoods))

	foods := []string{
		"Nasi Goreng",
		"Mie Ayam",
		"Ayam Goreng",
		"Bubur Ayam"}

	newFoods := foods[0:2]
	// [index, length]
	// [0,2] => mulai dari index ke-0 ambil dua data
	// [1:] => index ke-1 sepanjang array

	fmt.Printf("foods => %#v \n", foods)
	fmt.Printf("newFoods => %#v \n", newFoods)

	fmt.Printf("food len(%d), cap(%d) \n", len(foods), cap(foods))
	fmt.Printf("newFoods len(%d), cap(%d) \n", len(newFoods), cap(newFoods))
	// kapasitas(cap) diambil seluruh panjang slice bahkan melewati ketetapan panjang [0,2] walaupun nilai/value nya tidak diambil(len)

	// untuk mengubah nilai newfood tapi tidak mengubah nilai slice food maka harus mendobrak cap dari newFood agar melebihi cap dari food terlebih dahulu

	for index, value := range foods {
		fmt.Printf("index (%d), value (%s) \n", index, value)
	}
	// for index := range foods {
	// 	fmt.Printf("index (%d) \n", index)
	// }
	// for _,value := range foods {
	// 	fmt.Printf("value (%s) \n", value)
	// }
}
