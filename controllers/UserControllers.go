package controllers

import (
	"hr_manajemen/models"
	"hr_manajemen/repository"
	"hr_manajemen/service"

	"github.com/gin-gonic/gin"
)

func GetUser2(c *gin.Context) {
	result, err := service.GetUsers()
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, result)
	return
}

func GetUserByNIK(c *gin.Context) {
	var user models.User
	nik := c.Params.ByName("nik")
	err := repository.GetUserByNIK(&user, nik)
	if err != nil {
		c.JSON(200, gin.H{"message": "NIK not found "})
	} else {
		c.JSON(200, user)
	}
}
func InsertUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := service.CreateUser(user)
	if err != nil {
		c.JSON(200, gin.H{"message": err})
		return
	}
	c.JSON(200, gin.H{"message": "successed !"})
	return
}

func UpdateUser(c *gin.Context) {
	var user models.User
	nik := c.Params.ByName("nik")
	err := repository.GetUserByNIK(&user, nik)
	if err != nil {
		c.JSON(200, user)
	}
	c.BindJSON(&user)
	err = repository.UpdateUser(&user, nik)
	if err != nil {
		c.JSON(200, gin.H{"message": err})
		return
	}
	c.JSON(200, user)
}
func DeleteUser(c *gin.Context) {
	var user models.User
	nik := c.Params.ByName("nik")
	err := repository.DeleteUser(&user, nik)
	if err != nil {
		c.JSON(200, gin.H{"message": err})
		return
	}
	c.JSON(200, gin.H{"message": "deleted !"})
}

func LoginUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(200, gin.H{" message": err.Error()})
		return
	}
	pesan := service.ValidateLogin(user)
	c.JSON(200, gin.H{"message": pesan})
	return
}

func SignUpUser(c *gin.Context) {
	var user models.User
	err:= c.Bind(&user)
	if err != nil {
		c.JSON(200, gin.H{" message": err.Error()})
		return
	}
	err = service.UserSignUp(user)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "successed !"})
	return
}