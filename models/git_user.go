package models
// import "github.com/jinzhu/gorm"

type GitUser struct {
	// gorm.Model
	ID uint
	Username     string
	AvatarUrl       string 
	GithubUrl string
}