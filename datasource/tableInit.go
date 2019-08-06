package datasource

import (
	"../models"
)

func Createtable() {
	GetDB().AutoMigrate(
		&models.Wechat{},
		&models.User{},
		&models.Book{},
		&models.Article{},
	)
}