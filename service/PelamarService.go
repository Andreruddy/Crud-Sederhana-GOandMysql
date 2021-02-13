package service

import (
	"hr_manajemen/models"
	"hr_manajemen/repository"
	"hr_manajemen/utils"
)

func GetPelamars() ([]models.Pelamar, error) {
	var pelamar []models.Pelamar
	err := repository.GetAllPelamar(&pelamar)
	if err != nil {
		return nil, err
	}
	return pelamar, nil
}

func CreatePelamar(dataPelamar models.Pelamar) error {
	md5password := utils.GetMD5Hash(dataPelamar.Password)
	dataPelamar.Password = md5password
	err := repository.CreatePelamar(&dataPelamar)
	if err != nil {
		return err
	}
	return nil
}

func ValidatLogin(dataPelamar models.Pelamar) string {
	md5password := utils.GetMD5Hash(dataPelamar.Password)
	username := dataPelamar.Username
	pelamarLogin, err := repository.LoginPelamar(&dataPelamar)
	if err != nil {
		return err.Error()
	}

	if pelamarLogin.Username != username {
		return "username tidak cocok"
	}
	if pelamarLogin.Password != md5password {
		return "password salah !"
	}
	return "anda telah login"

}
