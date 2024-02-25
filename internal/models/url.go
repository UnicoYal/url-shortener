package models

import (
	"errors"
	"url-shortener/internal/storage/postgres"

	"github.com/jinzhu/gorm"
)

var database *gorm.DB

type URL struct {
	*gorm.Model
	Alias string `gorm:"text" json: "alias"`
	Url   string `gorm:"text" json: "url"`
}

func init() {
	postgres.Connect()
	database = postgres.GetDB()
	database.AutoMigrate(&URL{})
}

func SaveURL(url string, alias string) *URL {
	u := &URL{}
	u.Alias = alias
	u.Url = url

	database.NewRecord(u)
	database.Create(&u)

	return u
}

func GetUrl(alias string) (object *URL, err error) {
	var urlToFind URL
	// Second variant
	// query, err := database.DB().Prepare("SELECT * FROM url WHERE alias = ?")
	// if err != nil {
	// 	logger.Info("DB error: %s", err)
	// }
	// err = query.QueryRow(alias).Scan(&urlToFind)

	dbResult := database.Where("alias = ?", alias).Find(&urlToFind)
	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	return &urlToFind, nil
}

func DeleteUrl(alias string) *URL {
	var urlToDelete URL

	database.Where("alias=", alias).Delete(&urlToDelete)

	return &urlToDelete
}
