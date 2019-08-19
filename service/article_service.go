package service

import (
	"../models"
	"../repo"
	"strings"
	"github.com/spf13/cast"
	 _ "fmt"
	// "github.com/russross/blackfriday"
)

type ArticleService interface {
	GetArticleList(m map[string]interface{}) (result models.Result)
	SaveArticle(m map[string]interface{}) (result models.Result)
}
type articleService struct {
}

func NewArticleService() ArticleService {
	return &articleService{}
}

var articleRepo = repo.NewArticleRepository()
var articleCategoryRepo = repo.NewArticleCategoryRepository()

func (u articleService) GetArticleList(m map[string]interface{}) (result models.Result){
	result.Code = 0
	total,wechats := articleRepo.GetArticleList(m)
	maps := make(map[string]interface{},2)
	maps["Total"] = total
	maps["List"] = wechats
	result.Data = maps
	return
}
func (u articleService) SaveArticle(m map[string]interface{}) (result models.Result){
	result.Code = 0
	article,err := articleRepo.SaveArticle(m)
	if err != nil{
		result.Code = -1
		result.Msg = cast.ToString(err)
		return
	}
	articleID := article.ID;
	ids := strings.Split(cast.ToString(m["IDs"]),",")
	for _,i := range ids{
		articleCategoryRepo.SaveArticleCategory(articleID,cast.ToUint(i))
	}
	maps := make(map[string]interface{},1)
	maps["article"] = article
	result.Data = maps
	result.Msg = "保存成功"
	return
}


