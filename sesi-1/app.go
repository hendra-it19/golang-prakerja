package main

import "fmt"

func main() {

	var address string

	// variabel akan digunakan nanti)
	// _ = address

	// asign nilai ke variabel
	address = `Baubau`
	fmt.Println(address)

	// re asign nilai ke variabel
	address = `kendari`
	fmt.Println(address)

	var num1 uint8 = 10
	var num2 int8 = -10
	var num3 uint16 = 256
	var num4 float32 = 12.34
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)

	var b bool
	b = true
	fmt.Println(b)

	var name string = "Hendra"
	var result string = name + " Tinggal di " + address
	fmt.Println(result)

	var age uint8 = 34

	// concatenation
	fmt.Printf("%s tinggal di %s, umurnya %d tahun \n", name, address, age)
	fmt.Printf("Tipe data name adalah %T", name)
	// %t untuk boolean
}

// buatlah 3 buat variabel dengan tipe data string, int dan boolean
// kemudian cobalah print menggunakan fmt.Printf
