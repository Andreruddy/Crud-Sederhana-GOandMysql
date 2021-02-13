package repository

import (
	"hr_manajemen/config"
	"hr_manajemen/models"
	"hr_manajemen/utils"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllPelamar(pelamar *[]models.Pelamar) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Find(pelamar)
	return nil
}

func GetPelamarByNIK(pelamar *models.Pelamar, nik string) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Where("nik = ?", nik).First(pelamar)
	return nil
}

func CreatePelamar(pelamar *models.Pelamar) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Create(pelamar)
	return nil
}

func DeletePelamar(pelamar *models.Pelamar, nik string) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Where("nik = ?", nik).Delete(pelamar)
	return nil
}

func UpdatePelamar(pelamar *models.Pelamar) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Table("pelamar").Where("nik=?", pelamar.Nik).Update(pelamar)
	return nil
}

func UploadFoto(pelamar *models.Pelamar, foto string) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Create(foto).First(pelamar)
	return nil
}

func LoginPelamar(pelamar *models.Pelamar) (models.Pelamar, error) {
	var userPelamar models.Pelamar
	db, err := config.Connect()
	if err != nil {
		return userPelamar, err
	}
	defer db.Close()
	db.Table("pelamar").Where("username = ? AND password = ?", pelamar.Username, utils.GetMD5Hash(pelamar.Password)).First(&userPelamar)
	return userPelamar, nil
}

