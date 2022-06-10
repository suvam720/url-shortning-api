package models

import (
	"github.com/jinzhu/gorm"
)

type Url struct {
	gorm.Model
	LongUrl  string `json:"long_url" gorm:"unique"`
	ShortUrl string `json:"short_url" gorm:"unique"`
}
