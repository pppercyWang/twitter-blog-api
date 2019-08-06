package controllers

import (
	"github.com/kataras/iris"
	"log"
	"../models"
	"../service"
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
