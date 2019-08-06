package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Content        string `gorm:"type:text"`
	ArticleTypeId     uint `gorm:"default:0"`
	Title     string `gorm:"type:varchar(50);not null;"`
	Description     string `gorm:"type:varchar(255);"`
	Likes int `gorm:"default:0"`
}