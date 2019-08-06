package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Content        string `gorm:"size:5000"`
	Type     string `gorm:"type:varchar(20);not null;"`
	Title     string `gorm:"type:varchar(50);not null;"`
	Description     string `gorm:"type:varchar(255);"`
	Likes int `gorm:"default:0"`
}