package view

import (
	"fmt"
	controller "strukturdataMVC/controller/sustercontroller"
)

func InsertSuster() {
	var id int
	var nama, tlp, shift, tugas string
	fmt.Print("masukkan ID Baru : ")
	fmt.Scan(&id)
	fmt.Print("masukkan Nama Baru : ")
	fmt.Scan(&nama)

	err := controller.ControllerInsertSuster(id, nama, tlp, shift, tugas)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert berhasil")
	}
}

func UpdateDoter() {
	var id int
	var nama, tlp, shift, tugas string
	fmt.Print("masukkan ID : ")
	fmt.Scan(&id)
	fmt.Print("masukkan Nama Baru : ")
	fmt.Scan(&nama)
	err := controller.ControllerUpdate(id, nama, tlp, shift, tugas)
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
		fmt.Println(current.Next.Data.Id)
		fmt.Println(current.Next.Data.Nama)
	} else {
		fmt.Println("data tidak di temukan")
	}
}

func ViewSuster() {
	allSuster := controller.ControllerFindAllSuster()
	// fmt.Println(allSuster)
	for _, member := range allSuster {
		fmt.Println("id Suster :", member.Id)
		fmt.Println("Nama Suster :", member.Nama)
	}
}
