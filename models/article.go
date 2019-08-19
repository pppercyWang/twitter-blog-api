package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Content        string `gorm:"type:text"`
	Title     string `gorm:"type:varchar(80);not null;"`
	Description     string `gorm:"type:varchar(255);"`
	Likes uint `gorm:"default:0"`
	Personal uint `gorm:"default:0"`
}