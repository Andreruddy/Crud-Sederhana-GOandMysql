package config

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}


func Connect() (*gorm.DB, error) {
	return gorm.Open("mysql", "root:@/hr_manajemen")
}

