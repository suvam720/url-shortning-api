package helpers

import (
	"log"
	"math/rand"

	"github.com/suvam720/url-shortner/pkg/database"
	"github.com/suvam720/url-shortner/pkg/models"
	"github.com/suvam720/url-shortner/pkg/request"
)

func CreateUrl(req *request.RequestBody, random string) error {
	tx := database.Db.Create(&models.Url{
		LongUrl:  req.ReqUrl,
		ShortUrl: random,
	})

	return tx.Error
}

func FindUrl(url string) (u string, err error) {
	tx := database.Db.Table("urls").Select("long_url").Where("short_url = ?", url).Scan(&u)
	if tx.Error != nil {
		return "", tx.Error
	}

	return
}

func GetAllUrl() ([]models.Url, error) {
	var urls []models.Url
	tx := database.Db.Find(&urls)
	if tx.Error != nil {
		return urls, tx.Error
	}

	return urls, nil
}

func DeleteUrl(url string) error {
	durl, err := FindUrl(url)
	if err != nil {
		log.Fatal(err)
	}
	tx := database.Db.Delete(&durl)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

var runes = []rune("abcdefghijklmnopqrstuvwzyx!@#$%&0123456789")

func Random(size int) string {
	str := make([]rune, size)
	for i := range str {
		str[i] = runes[rand.Intn(len(runes))]
	}

	return string(str)
}
