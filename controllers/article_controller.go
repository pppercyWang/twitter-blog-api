package controllers

import (
	"github.com/kataras/iris"
	"log"
	"../models"
	"../service"
	// "fmt"
	"strings"
	"github.com/spf13/cast"
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
	if m["Page"] == "" || m["Page"] == nil {
		result.Code = -1
		result.Msg = "参数缺失 Page"
		return
	}
	if cast.ToUint(m["Page"]) == 0 {
		result.Code = -1
		result.Msg = "参数错误 Page"
		return
	}
	if m["Size"] == "" || m["Size"] == nil {
		result.Code = -1
		result.Msg = "参数缺失 Size"
		return
	}
	if cast.ToUint(m["Size"]) == 0 {
		result.Code = -1
		result.Msg = "参数错误 Size"
		return
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
	if m["CategoryIDs"] == "" || m["CategoryIDs"] == nil {
		result.Code = -1
		result.Msg = "请选择至少一个文章分类"
		return
	}
	ids := strings.Split(cast.ToString(m["CategoryIDs"]),",")
	for _,i := range ids {
		if cast.ToUint(i) == 0 {
			result.Code = -1
			result.Msg = "参数错误 CategoryIDs"
			return
		}
	}
	return g.Service.SaveArticle(m)
}

func (g *ArticleController) PostFetch() (result models.Result)  {
	var m map[string]interface{}
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	if m["ID"] == "" || m["ID"] == nil {
		result.Code = -1
		result.Msg = "参数错误 ID"
		return
	}
	return g.Service.GetArticle(m)
}