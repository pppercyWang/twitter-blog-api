/*
@Time : 2019/4/8 10:03
@Author : lukebryan
@File : user_repo
@Software: GoLand
*/
package repo
import (
	"../datasource"
	"../models"
	// "github.com/spf13/cast"
)
type GitUserRepository interface {
	SaveGitUserInfo(gitUser models.GitUser) ()
}
func NewGitUserRepository() GitUserRepository {
	return &gitUserRepository{}
}
type gitUserRepository struct{}
func (n gitUserRepository) SaveGitUserInfo(gitUser models.GitUser) () {
	db := datasource.GetDB()
	db.Save(&gitUser);
}
