/*
@Time : 2019/4/3 11:09 
@Author : lukebryan
@File : NoticeService
@Software: GoLand
*/
package service

import (
	"../models"
	"../repo"
)

type WechatService interface {
	List(m map[string]interface{}) (result models.Result)
}
type wechatServices struct {

}

func NewWechatServices() WechatService {
	return &wechatServices{}
}

var wechatRepo = repo.NewWechatRepository()


func (u wechatServices) List(m map[string]interface{}) (result models.Result){
	result.Code = 0
	total,wechats := wechatRepo.List(m)
	maps := make(map[string]interface{},2)
	maps["Total"] = total
	maps["List"] = wechats
	result.Data = maps
	return
}
