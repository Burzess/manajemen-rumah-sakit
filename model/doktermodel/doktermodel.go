package doktermodel

import (
	DB "strukturdataMVC/Database"
	entity "strukturdataMVC/entities"
)

func ModelInsertDokter(container entity.Dokter) {
	newGerbong := entity.LinkedlistDokter{}
	newGerbong.Data = container
	temp := &DB.DBDokter
	if temp.Next == nil {
		temp.Next = &newGerbong
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &newGerbong
	}
}

func Search(id int) *entity.LinkedlistDokter {
	current := DB.DBDokter.Next
	for current != nil {
		if current.Data.Id == id {
			return current
		}
		current = current.Next
	}
	return nil
}

func ModelUpdateDokter(newData entity.Dokter) bool {
	current := &DB.DBDokter
	for current != nil {
		if current.Data.Id == newData.Id {
			current.Data = newData
			return true
		}
		current = current.Next
	}
	return false
}

func ModelDeleteDokter(id int) bool {
	prev := &DB.DBDokter
	current := DB.DBDokter.Next

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

func ModelViewAllDokter() []entity.Dokter {
	temp := DB.DBDokter.Next
	members := []entity.Dokter{}
	for temp != nil {
		members = append(members, temp.Data)
		temp = temp.Next
	}
	return members
}
