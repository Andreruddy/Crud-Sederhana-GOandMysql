package repository

import (
	"hr_manajemen/config"
	"hr_manajemen/models"

	_ "github.com/go-sql-driver/mysql"
)

func InsertPerusahaan(perusahaan *models.Perusahaan) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Create(perusahaan)
	return nil
}

func GetAllPerusahaan(perusahaan *[]models.Perusahaan) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Find(perusahaan)
	return nil
}

func GetPerusahaanByID(perusahaan *models.Perusahaan, id string) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Where("id = ?", id).First(perusahaan)
	return nil
}

func DeletePerusahaan(perusahaan *models.Perusahaan, id string) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Where("id = ?", id).Delete(perusahaan)
	return nil
}

func UpdatePerusahaan(perusahaan *models.Perusahaan) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Table("perusahaan").Update(perusahaan)
	return nil
}