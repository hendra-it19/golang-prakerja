package main

import (
	"fmt"
	"os"
)

type Student struct {
	Id         string
	Name       string
	Address    string
	Occupation string
	Reason     string
}

var students = []Student{
	{
		Id:         "1",
		Name:       "Hendra",
		Address:    "Kendari",
		Occupation: "Programmer",
		Reason:     "Belajar",
	},
	{
		Id:         "2",
		Name:       "Dian",
		Address:    "Baubau",
		Occupation: "Mahasiswa",
		Reason:     "Ingin jadi be developer",
	},
	{
		Id:         "3",
		Name:       "Indra",
		Address:    "Baubau",
		Occupation: "Mahasiswa",
		Reason:     "Ingin jadi be developer",
	},
	{
		Id:         "4",
		Name:       "Citra",
		Address:    "Kolaka",
		Occupation: "Backend Engineer",
		Reason:     "Ingin jadi be golang",
	},
	{
		Id:         "5",
		Name:       "Gita",
		Address:    "Kendari",
		Occupation: "Backend Engineer",
		Reason:     "Ingin menambah bahasa pemrograman baru",
	},
}

// method
func (value Student) PrintDetailData() {
	fmt.Printf("Nama %s, Alamat %s, Pekerjaan %s, Alasan %s", value.Name, value.Address, value.Occupation, value.Reason)
}

func findStudent(id string) {
	for _, value := range students {
		if value.Id == id {
			// memanggil method
			value.PrintDetailData()
			return
		}
	}
	fmt.Printf("Siswa dengan ID %s, tidak ditemukan !", id)
}

func main() {
	var args []string = os.Args
	// fmt.Printf("%#v \n", args[1])
	if len(args) < 2 {
		fmt.Println("Id siswa wajib diisi!")
		return
	}
	findStudent(args[1])
}
