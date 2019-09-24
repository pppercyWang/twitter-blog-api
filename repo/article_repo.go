/*
@Time : 2019/4/8 10:03
@Author : lukebryan
@File : user_repo
@Software: GoLand
*/
package repo

import (
	"../datasource"
	"../models"
	// "../util"
	"github.com/spf13/cast"
	// "log"
	// "fmt"
	// "strings"
	// "reflect"
	// "io/ioutil"
	// "testing"
)

type ArticleRepository interface {
	GetArticleList(m map[string]interface{}) (total int, articles []models.Article)
	SaveArticle(m map[string]interface{}) (article models.Article,err error)
}

func NewArticleRepository() ArticleRepository {
	return &articleRepository{}
}

type articleRepository struct{}
func (n articleRepository)GetArticleList(m map[string]interface{})(total int,articles []models.Article){
	db := datasource.GetDB()
	err := db.Limit(cast.ToInt(m["Size"])).Offset((cast.ToInt(m["Page"])-1)*cast.ToInt(m["Size"])).Find(&articles).Error
	if err!=nil {
		panic("select Error")
	}
	datasource.GetDB().Model(&models.Article{}).Count(&total)
	return
}
func (n articleRepository)SaveArticle(m map[string]interface{})(article models.Article,err error){
	content := cast.ToString(m["Content"])
	article.Content = content
	article.Title = cast.ToString(m["Title"])
	article.Personal = cast.ToUint(m["Personal"])
	article.Tags = cast.ToString(m["Tags"])
	db := datasource.GetDB()
	err = db.Save(&article).Error;
	return
}
