package controllers

import (
	"github.com/kataras/iris"
	"log"
	"../models"
	"../service"
	// "fmt"
)

type ArticleController struct {
	Ctx     iris.Context
	Service service.ArticleService
}

func NewArticleController() *ArticleController {
	return &ArticleController{ Service: service.NewArticleService() }
}

func (g *ArticleController) PostList() (result models.Result)  {
	var m map[string]interface{}
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	return g.Service.GetArticleList(m)
}
func (g *ArticleController) PostSave() (result models.Result)  {
	var m map[string]interface{}
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	if m["Title"] == "" || m["Title"] == nil {
		result.Code = -1
		result.Msg = "请输入文章标题"
		return
	}
	if m["Content"] == "" || m["Content"] == nil {
		result.Code = -1
		result.Msg = "请输入文章内容"
		return
	}
	if m["Personl"] == "" || m["Personal"] == nil {
		result.Code = -1
		result.Msg = "请选择文章类型"
		return
	}
	if m["IDs"] == "" || m["IDs"] == nil {
		result.Code = -1
		result.Msg = "请选择至少一个文章分类"
		return
	}
	return g.Service.SaveArticle(m)
}

