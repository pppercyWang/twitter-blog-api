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
	GetArticle(articleID uint) (article models.Article)
	SaveArticle(m map[string]interface{}) (article models.Article,err error)
	SaveArticleCategories(articleID uint,categoryStr string)(article models.Article)
}

func NewArticleRepository() ArticleRepository {
	return &articleRepository{}
}

type articleRepository struct{}
func (n articleRepository)GetArticle(articleID uint)(article models.Article){
	db := datasource.GetDB()
	db.First(&article, articleID)
	return
}
func (n articleRepository)GetArticleList(m map[string]interface{})(total int,articles []models.Article){
	db := datasource.GetDB()
	var  err error; 
	if m["Personal"] == nil {

		datasource.GetDB().Model(&models.Article{}).Count(&total)
		 err = db.Limit(cast.ToInt(m["Size"])).Offset((cast.ToInt(m["Page"])-1)*cast.ToInt(m["Size"])).Order("created_at desc").Find(&articles).Error
	}else{
		datasource.GetDB().Model(&models.Article{}).Where("personal = ?", cast.ToInt(m["Personal"])).Count(&total)
		err =  db.Limit(cast.ToInt(m["Size"])).Offset((cast.ToInt(m["Page"])-1)*cast.ToInt(m["Size"])).Where("personal = ?", cast.ToInt(m["Personal"])).Order("created_at desc").Find(&articles).Error
	}
	if err!=nil {
		panic("select Error")
	}
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
func (n articleRepository)SaveArticleCategories(articleID uint,categoryStr string)(article models.Article){
	article.ID = articleID
	db := datasource.GetDB()
	db.Model(&article).Update("categories", categoryStr)
	return
}
