package database

import (
	"github.com/suvam720/url-shortner/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect(dsn string) error {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate(&models.Url{})
	Db = db

	return nil
}
