package service

// menggunakan huruf besar agar bisa diakses globa
// menggunakan huruf kecil akan menjadi private dan hanya bisa di akses di dalam pakcage itu sendiri
type Product struct {
	Title    string
	Price    int
	Quantity int
}
