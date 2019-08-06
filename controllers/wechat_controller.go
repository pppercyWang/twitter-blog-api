/*
@Time : 2019/4/2 17:18 
@Author : lukebryan
@File : DeviceController
@Software: GoLand
*/
package controllers

import (
	"github.com/kataras/iris"
	"log"
	"../models"
	"../service"
)

type WechatController struct {
	Ctx     iris.Context
	Service service.WechatService
}

func NewWechatController() *WechatController {
	return &WechatController{ Service: service.NewWechatServices() }
}

func (g *WechatController) PostList() (result models.Result)  {
	var m map[string]interface{}
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	return g.Service.List(m)
}
