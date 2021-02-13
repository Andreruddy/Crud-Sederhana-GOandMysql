package repository

import (
	"fmt"
	"hr_manajemen/config"
	"hr_manajemen/models"
	"hr_manajemen/utils"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllUser(user *[]models.User) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Find(user)
	return nil
}

func GetLogin(user *models.User, username string) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Where("username = ?", username).First(user)
	return nil
}

func GetUserByNIK(user *models.User, nik string) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Where("nik = ?", nik).First(user)
	return nil
}

func CreateUser(user *models.User) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Create(user)
	return nil
}

func UpdateUser(user *models.User, nik string) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	fmt.Println(user)
	db.Save(user)
	return nil
}

func DeleteUser(user *models.User, nik string) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Where("nik = ?", nik).Delete(user)
	return nil
}

func UserLogin(user *models.User) (models.User, error) {
	var dataUser2 models.User
	db, err := config.Connect()
	if err != nil {
		return dataUser2, err
	}
	defer db.Close()
	db.Where("username = ? AND password = ?", user.Username, utils.GetMD5Hash(user.Password)).First(&dataUser2)
	return dataUser2, nil
}

func SignUp(user *models.User) (err error) {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Create(user)
	return nil
}
