package service

import (
	"../models"
	// "../repo"
	// "fmt"
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
	return
}

