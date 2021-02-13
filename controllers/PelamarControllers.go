package controllers

import (
	// "hr_manajemen/model
	"fmt"
	"hr_manajemen/models"
	"hr_manajemen/repository"
	"hr_manajemen/service"
	"hr_manajemen/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPelamar(c *gin.Context, req http.Request) {
	limit := 5
	page, begin := utils.Paginate(&req, limit)
	fmt.Printf("Current Page: %d, Begin %d \n", page, begin )
	result, err := service.GetPelamars()
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, result)
}
func GetPelamarsByNIK(c *gin.Context) {
	var pelamar models.Pelamar
	nik := c.Params.ByName("nik")
	err := repository.GetPelamarByNIK(&pelamar, nik)
	if err != nil {
		c.JSON(200, gin.H{"message": "Nik tidak ditemukan"})
		return
	}
	c.JSON(200, pelamar)
}

func DeletePelamars(c *gin.Context) {
	var pelamar models.Pelamar
	nik := c.Params.ByName("nik")
	err := repository.DeletePelamar(&pelamar, nik)
	if err != nil {
		c.JSON(200, gin.H{"message": err})
		return
	}
	c.JSON(200, gin.H{"message": "deleted !"})
}

func UpdatePelamars(c *gin.Context) {
	var pelamar models.Pelamar
	err := c.Bind(&pelamar)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	err = repository.UpdatePelamar(&pelamar)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "data Updated !"})
}

func UploadFoto(c *gin.Context) {
	foto, err := c.FormFile("file")
	fmt.Println(foto)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	path := "images/" + foto.Filename
	if err := c.SaveUploadedFile(foto, path); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", foto.Filename))
}

func InsertPelamar(c *gin.Context) {
	var pelamar models.Pelamar
	err := c.BindJSON(&pelamar)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	err = service.CreatePelamar(pelamar)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "insert success !"})
	return
}

func LoginPelamar(c *gin.Context) {
	var pelamar models.Pelamar
	err := c.BindJSON(&pelamar)
	if err != nil {
		c.JSON(200, gin.H{" message": err.Error()})
		return
	}
	pesan := service.ValidatLogin(pelamar)
	c.JSON(200, gin.H{"message": pesan})
	return
}

