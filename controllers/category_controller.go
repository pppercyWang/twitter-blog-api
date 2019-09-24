package controllers

import (
	"github.com/kataras/iris"
	"log"
	"../models"
	"../service"
	// "fmt"
)

type CategoryController struct {
	Ctx     iris.Context
	Service service.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{ Service: service.NewCategoryService() }
}
func (g *CategoryController) PostSave() (result models.Result)  {
	var m map[string]interface{}
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	if m["Name"] == "" || m["Name"] == nil {
		result.Code = -1
		result.Msg = "请输入分类名称"
		return
	}
	return g.Service.SaveCategory(m)
}
func (g *CategoryController) PostList() (result models.Result)  {
	var m map[string]interface{}
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	return g.Service.GetCategoryList(m)
}

