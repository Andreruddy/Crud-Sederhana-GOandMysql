package service

import (
	"hr_manajemen/models"
	"hr_manajemen/repository"
)
func CreatePerusahaan(perusahaan models.Perusahaan) error {
	err := repository.InsertPerusahaan(&perusahaan)
	if err != nil {
		return err
	}
	return nil
}

func GetAllPerusahaan() ([]models.Perusahaan, error) {
	var perusahaan []models.Perusahaan
	err := repository.GetAllPerusahaan(&perusahaan)
	if err != nil {
		return nil, err
	}
	return perusahaan, nil
}