package entities

type Dokter struct {
	Id        int
	Nama      string
	Tlp       string
	JamKerja  string
	Spesialis string
}
type LinkedlistDokter struct {
	Data Dokter
	Next *LinkedlistDokter
}

type Suster struct {
	Id    int
	Nama  string
	Tlp   string
	Shift string
	Tugas string
}
type LinkedlistSuster struct {
	Data Suster
	Next *LinkedlistSuster
}

type Medis struct {
	Tanggal string
	Peyakit string
}
type Pasien struct {
	Id      int
	Nama    string
	Tlp     string
	Kondisi Medis
}
type LinkedlistPasien struct {
	Data Pasien
	Next *LinkedlistPasien
}
