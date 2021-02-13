package service

import (
	"errors"
	"hr_manajemen/models"
	"hr_manajemen/repository"
	"hr_manajemen/utils"
)

func GetUsers() ([]models.User, error) {
	var user []models.User
	err := repository.GetAllUser(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(dataUser models.User) error {
	md5password := utils.GetMD5Hash(dataUser.Password)
	dataUser.Password = md5password
	err := repository.CreateUser(&dataUser)
	if err != nil {
		return err
	}
	return nil
}

func ValidateLogin(dataUser models.User) string {
	md5password := utils.GetMD5Hash(dataUser.Password)
	username := dataUser.Username
	userLogin, err := repository.UserLogin(&dataUser)
	if err != nil {
		return err.Error()
	}

	if userLogin.Username != username {
		return "username tidak cocok"
	}
	if userLogin.Password != md5password {
		return "password salah !"
	}
	return "anda telah login"

}
func UserSignUp(dataUser models.User) error {
	if dataUser.Username == "" {
		return errors.New("User tidak boleh kosong")
	}
	if dataUser.Nama == "" {
		return errors.New("Nama tidak boleh kosong")
	}
	if dataUser.Email == "" {
		return errors.New("Email tidak boleh kosong")
	}
	if dataUser.Password == "" {
		return errors.New("Password tidak boleh kosong")
	}
	md5password := utils.GetMD5Hash(dataUser.Password)
	dataUser.Password = md5password
	err := repository.SignUp(&dataUser)
	if err != nil {
		return err
	}
	return nil
}