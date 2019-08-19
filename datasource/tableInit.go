package datasource

import (
	"../models"
)

func Createtable() {
	GetDB().AutoMigrate(
		&models.User{},
		&models.Article{},
		&models.Category{},
		&models.ArticleCategory{},
	)
}