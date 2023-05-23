package dokterview

import (
	"fmt"
	controller "strukturdataMVC/controller/doktercontroller"
)

func InsertDokter() {
	var id int
	var nama, tlp, jamKerja, spesialis string
	fmt.Print("masukkan ID Baru : ")
	fmt.Scan(&id)
	fmt.Print("masukkan Nama Baru : ")
	fmt.Scan(&nama)
	// fmt.Print("masukkan Telepon Baru : ")
	// fmt.Scan(&tlp)
	// fmt.Print("masukkan Jam Kerja Baru : ")
	// fmt.Scan(&jamKerja)
	// fmt.Print("masukkan Spesialis Baru : ")
	// fmt.Scan(&spesialis)

	err := controller.ControllerInsertDokter(id, nama, tlp, jamKerja, spesialis)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert berhasil")
	}
}

func UpdateDoter() {
	var id int
	var nama, tlp, jamKerja, spesialis string
	fmt.Print("masukkan ID : ")
	fmt.Scan(&id)
	fmt.Print("masukkan Nama Baru : ")
	fmt.Scan(&nama)
	// fmt.Print("masukkan Telepon Baru : ")
	// fmt.Scan(&tlp)
	// fmt.Print("masukkan Jam Kerja Baru : ")
	// fmt.Scan(&jamKerja)
	// fmt.Print("masukkan Spesialis Baru : ")
	// fmt.Scan(&spesialis)
	err := controller.ControllerUpdate(id, nama, tlp, jamKerja, spesialis)
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

func ViewDokter() {
	allDokter := controller.ControllerFindAllDokter()
	// fmt.Println(allDokter)
	for _, member := range allDokter {
		fmt.Println("id Dokter :", member.Id)
		fmt.Println("Nama Dokter :", member.Nama)
	}
}
