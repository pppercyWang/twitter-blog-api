package service

import (
	"../models"
	"../repo"
	// "github.com/russross/blackfriday"
	// "fmt"
)

type CategoryService interface {
	SaveCategory(m map[string]interface{}) (result models.Result)
	GetCategoryList(m map[string]interface{}) (result models.Result)
}
type categoryService struct {
}

func NewCategoryService() CategoryService {
	return &categoryService{}
}

var categoryRepo = repo.NewCategoryRepository()


func (u categoryService) SaveCategory(m map[string]interface{}) (result models.Result){
	if m["Name"] == "" || m["Name"] == nil {
		result.Code = -1
		result.Msg = "请输入文章类型名称！"
		return
	}
	result.Code = 0
	category := categoryRepo.SaveCategory(m)
	maps := make(map[string]interface{},1)
	maps["category"] = category
	result.Data = maps
	result.Msg = "保存成功"
	return
}
func (u categoryService) GetCategoryList(m map[string]interface{}) (result models.Result){
	categorys := categoryRepo.GetCategoryList(m)
	result.Code = 0
	maps := make(map[string]interface{},1)
	maps["List"] = categorys
	result.Data = maps
	return
}


