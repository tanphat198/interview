package config

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	gormDB, _ := gorm.Open("sqlite3", "manage.db")
	return gormDB
}
