package pasienmodel

import (
	DB "strukturdataMVC/Database"
	entity "strukturdataMVC/entities"
)

func ModelInsertPasien(container entity.Pasien) {
	newGerbong := entity.LinkedlistPasien{}
	newGerbong.Data = container
	temp := &DB.DBPasien
	if temp.Next == nil {
		temp.Next = &newGerbong
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &newGerbong
	}
}

func Search(id int) *entity.LinkedlistPasien {
	current := DB.DBPasien.Next
	for current != nil {
		if current.Data.Id == id {
			return current
		}
		current = current.Next
	}
	return nil
}

func ModelUpdatePasien(newData entity.Pasien) bool {
	current := &DB.DBPasien
	for current != nil {
		if current.Data.Id == newData.Id {
			current.Data = newData
			return true
		}
		current = current.Next
	}
	return false
}

func ModelDeletePasien(id int) bool {
	prev := &DB.DBPasien
	current := DB.DBPasien.Next

	for current != nil {
		if current.Data.Id == id {
			prev.Next = current.Next
			return true
		}
		prev = current
		current = current.Next
	}

	return false
}

func ModelViewAllPasien() []entity.Pasien {
	temp := DB.DBPasien.Next
	members := []entity.Pasien{}
	for temp != nil {
		members = append(members, temp.Data)
		temp = temp.Next
	}
	return members
}
