package view

import "fmt"

func MenuUtama() {
	println("1. Dokter")
	println("2. Suster")
	println("3. Pasien")
	println("4. Exit")
	fmt.Print("Pilih: ")
}
func SubMenu() {
	fmt.Println("1. Insert")
	fmt.Println("2. Update")
	fmt.Println("3. Delete")
	fmt.Println("4. View By Id")
	fmt.Println("5. View All")
	fmt.Println("6. Exit")
	fmt.Print("Pilih: ")
}
