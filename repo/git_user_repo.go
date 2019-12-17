
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
