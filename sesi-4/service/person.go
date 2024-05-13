package service

// menggunakan huruf besar agar bisa diakses global
// menggunakan huruf kecil akan menjadi private dan hanya bisa di akses di dalam pakcage itu sendiri
type Person struct {
	Name    string
	Age     int
	Role    string
	address string //dibuat private
}

// role
// fungsi agar dapat mengakses nilai variabel private di luar package
func (p Person) Address() string {
	if p.Role == "admin" {
		return p.address
	}
	return "Access denied!"
}

// fungsi untuk mengisi variabel private dan global di luar package
func NewPerson(name string, age int, role string, address string) Person {
	// assign nilai parameter fungsi kedalam struct Person
	return Person{
		Name:    name,
		Age:     age,
		address: address,
		Role:    role,
	}
}
