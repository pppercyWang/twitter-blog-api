package models
import (
	"github.com/jinzhu/gorm"
)
type Comment struct {
	gorm.Model
	Content string
	ArticleID uint
	GitUserID uint
	Username     string
	AvatarUrl       string 
	GithubUrl string
	ArticleTitle string
}