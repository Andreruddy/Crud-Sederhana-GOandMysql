package models

type Perusahaan struct {
	ID             int    `json:"id"`
	NamaPerusahaan string `json:"nama_perusahaan"`
	Alamat         string `json:"alamat"`
	NoTelephone    string `json:"no_telephone"`
}

func (b *Perusahaan) TableName() string {
	return "Perusahaan"
}
