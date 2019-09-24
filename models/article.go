package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Content        string `gorm:"type:text"`
	Title     string `gorm:"type:varchar(80);not null;"`
	Tags  string `gorm:"type:varchar(20)"`
	Personal uint `gorm:"default:0"`	
}