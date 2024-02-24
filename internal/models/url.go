package models

import (
	"url-shortener/internal/storage/postgres"

	"github.com/jinzhu/gorm"
)

var database *gorm.DB

type URL struct {
	*gorm.Model
	Alias string `gorm:""json: "alias"`
	Url   string `json: "url"`
}

func init() {
	postgres.Connect()
	database = postgres.GetDB()
	database.AutoMigrate(&URL{})
}

func (u *URL) SaveURL(url string, alias string) *URL {
	u.Alias = alias
	u.Url = url

	database.NewRecord(u)
	database.Create(&u)

	return u
}

func GetUrl(alias string) *URL {
	var urlToFind *URL
	// Second variant
	// query, err := database.DB().Prepare("SELECT * FROM url WHERE alias = ?")
	// if err != nil {
	// 	logger.Info("DB error: %s", err)
	// }
	// err = query.QueryRow(alias).Scan(&urlToFind)

	database.Where("alias=", alias).Find(&urlToFind)

	return urlToFind
}

func DeleteUrl(alias string) *URL {
	var urlToDelete *URL

	database.Where("alias=", alias).Delete(&urlToDelete)

	return urlToDelete
}
