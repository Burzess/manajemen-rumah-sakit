package doktercontroller

import (
	"errors"
	"strukturdataMVC/entities"
	model "strukturdataMVC/model/doktermodel"
)

func ControllerInsertDokter(id int, nama, tlp, jamKerja, spesialis string) error {
	find := model.Search(id)
	if find == nil {
		data := entities.Dokter{
			Id:        id,
			Nama:      nama,
			Tlp:       tlp,
			JamKerja:  jamKerja,
			Spesialis: spesialis,
		}
		model.ModelInsertDokter(data)
		return nil
	}

	return errors.New("insert Data gagal, id yang di inputkan sudah tersedia")
}

func ControllerUpdate(id int, nama, tlp, jamKerja, spesialis string) error {
	container := entities.Dokter{
		Id:        id,
		Nama:      nama,
		Tlp:       tlp,
		JamKerja:  jamKerja,
		Spesialis: spesialis,
	}

	if model.ModelUpdateDokter(container) {
		return nil
	}

	return errors.New("update gagal id tidak ditemukan")
}

func ControllerDelete(id int) error {
	err := model.ModelDeleteDokter(id)
	if err {
		return nil
	}

	return errors.New("delete gagal id tidak ditemukan")
}

func ControllerViewById(id int) *entities.LinkedlistDokter {
	current := model.Search(id)
	if current == nil {
		return nil
	}
	return current
}

func ControllerFindAllDokter() []entities.Dokter {
	return model.ModelViewAllDokter()
}
