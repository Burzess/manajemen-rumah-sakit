package view

import (
	"fmt"
	controller "strukturdataMVC/controller/pasiencontroller"
)

func InsertPasien() {
	var id int
	var nama, tlp, tgl, penyakit string
	fmt.Print("masukkan ID Baru : ")
	fmt.Scan(&id)
	fmt.Print("masukkan Nama Baru : ")
	fmt.Scan(&nama)

	err := controller.ControllerInsertPasien(id, nama, tlp, tgl, penyakit)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert berhasil")
	}
}

func UpdateDoter() {
	var id int
	var nama, tlp, tgl, penyakit string
	fmt.Print("masukkan ID : ")
	fmt.Scan(&id)
	fmt.Print("masukkan Nama Baru : ")
	fmt.Scan(&nama)
	err := controller.ControllerUpdate(id, nama, tlp, tgl, penyakit)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("update berhasil")
	}
}

func DeleteDoter() {
	var id int
	fmt.Println("masukan ID")
	fmt.Scan(&id)
	err := controller.ControllerDelete(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("delete berhasil")
	}
}

func ViewById() {
	var id int
	fmt.Println("masukan ID")
	fmt.Scan(&id)
	current := controller.ControllerViewById(id)
	if current != nil {
		fmt.Println(current.Id)
		fmt.Println(current.Nama)
	} else {
		fmt.Println("data tidak di temukan")
	}
}

func ViewPasien() {
	allPasien := controller.ControllerFindAllPasien()
	// fmt.Println(allPasien)
	for _, member := range allPasien {
		fmt.Println("id Pasien :", member.Id)
		fmt.Println("Nama Pasien :", member.Nama)
	}
}
