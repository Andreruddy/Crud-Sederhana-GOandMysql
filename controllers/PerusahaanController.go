package controllers

import (
	"hr_manajemen/models"
	"hr_manajemen/repository"
	"hr_manajemen/service"

	"github.com/gin-gonic/gin"
)


func InsertPerusahaan(c *gin.Context) {
	var perusahaan models.Perusahaan
	c.BindJSON(&perusahaan)
	err := service.CreatePerusahaan(perusahaan)
	if err != nil {
		c.JSON(200, gin.H{"message": err})
		return
	}
	c.JSON(200, gin.H{"message": "successed !"})
	return
}

func GetAllPerusahaans(c *gin.Context) {
	result, err := service.GetAllPerusahaan()
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, result)
	return
}

func GetPerusahaanByID(c *gin.Context) {
	var perusahaan models.Perusahaan
	id := c.Params.ByName("id")
	err := repository.GetPerusahaanByID(&perusahaan, id)
	if err != nil {
	c.JSON(200, gin.H{"message":"ID Not Found"})
	}else {
		c.JSON(200, perusahaan)
	}
}

func DeletePerusahaan(c *gin.Context) {
	var perusahaan models.Perusahaan
	id := c.Params.ByName("id")
	err := repository.DeletePerusahaan(&perusahaan, id)
	if err != nil {
		c.JSON(200, gin.H{"message": err})
		return
	}
	c.JSON(200, gin.H{"message": "deleted !"})
}

func UpdatePerusahaan(c *gin.Context) {
	var perusahaan models.Perusahaan
	err := c.Bind(&perusahaan)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	err = repository.UpdatePerusahaan(&perusahaan)
	if err != nil {
		c.JSON(200, gin.H{"message": err})
		return
	}
	c.JSON(200,gin.H{"message":"data Updated !"})
}