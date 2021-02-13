package models

type Pelamar struct {
	Nik                int    `json:"nik"`
	Nama               string `json:"nama"`
	Username           string `json:"username"`
	Password           string `json:"password`
	Email              string `json:"email"`
	TempatTanggalLahir string `json:"tempat_tanggal_lahir"`
	Alamat             string `json:"alamat"`
	NoHandphone        string `json:"no_handphone"`
	Foto               string `json:"foto"`
	Status             string `json:"status"`
	Skill              string `json:"skill"`
	PendidikanTerakhir string `json:"pendidikan_terakhir"`
	JenisKelamin       string `json:"jenis_kelamin"`
}

func (b *Pelamar) TableName() string {
	return "pelamar"
}
