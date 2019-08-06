package service

import (
	"../models"
	"../repo"
	// "github.com/russross/blackfriday"
)

type ArticleService interface {
	GetArticleList(m map[string]interface{}) (result models.Result)
}
type articleService struct {
}

func NewArticleService() ArticleService {
	return &articleService{}
}

var articleRepo = repo.NewArticleRepository()


func (u articleService) GetArticleList(m map[string]interface{}) (result models.Result){
	result.Code = 0
	total,wechats := articleRepo.GetArticleList(m)
	maps := make(map[string]interface{},2)
	maps["Total"] = total
	maps["List"] = wechats
	result.Data = maps
	return
}
