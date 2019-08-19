package models

import (
	"github.com/jinzhu/gorm"
)

type ArticleCategory struct {
	gorm.Model
	CategoryID uint `gorm:"not null"`
	ArticleID uint `gorm:"not null"`
}