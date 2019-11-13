package service

import (
	"../models"
	// "../repo"
	// "fmt"
	"encoding/json"
	// "github.com/russross/blackfriday"
	"../util"
)
// https://github.com/login/oauth/access_token"
// https://api.github.com/user
type AuthService interface {
	GetUserInfo(m map[string]interface{}) (result models.Result)
}
type authService struct {
}
func NewAuthService() AuthService {
	return &authService{}
}
var gitUserService = NewGitUserServices()

type GitUserStruct struct {
    Login              string   `json:"login"`
    ID              uint      `json:"id"`
    AvatarUrl     string   `json:"avatar_url"`
    HtmlUrl string      `json:"html_url"`
}

func (u authService) GetUserInfo(m map[string]interface{}) (result models.Result){
	result.Code = 0
	retStr :=utils.PostRequest("https://github.com/login/oauth/access_token",m)
	token:= utils.GetUrlParam(retStr,"access_token")
	if token == "" {
		result.Code = -1
		result.Msg = "获取access_token失败"
	}
	retStr2 := utils.GetRequest("https://api.github.com/user?access_token=" + token)
	result.Data = retStr2
	var gitUserStruct GitUserStruct
	if err := json.Unmarshal([]byte(retStr2), &gitUserStruct); err == nil {
		var gitUser  = models.GitUser{Username:gitUserStruct.Login,AvatarUrl:gitUserStruct.AvatarUrl,ID:gitUserStruct.ID,GithubUrl: gitUserStruct.HtmlUrl}
		gitUserService.SaveGitUserInfo(gitUser)
    }
	return
}

