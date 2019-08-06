package models

import (
	"github.com/jinzhu/gorm"
)

type ArticleType struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null;"`
	Count     uint `gorm:"defualt:0"`
}