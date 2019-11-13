package service
import (
	"../models"
	"../repo"
	// "fmt"
	// "github.com/spf13/cast"
	// "log"
)
type GitUserService interface {
	SaveGitUserInfo(gitUser models.GitUser) (result models.Result)
}
type gitUserServices struct {}
func NewGitUserServices() GitUserService {
	return &gitUserServices{}
}
var gitUserRepo = repo.NewGitUserRepository()
/*
保存
 */
func (u gitUserServices) SaveGitUserInfo(gitUser models.GitUser) (result models.Result){
	gitUserRepo.SaveGitUserInfo(gitUser)
	return
}

