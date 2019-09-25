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
	// "github.com/spf13/cast"
	// "log"
	// "fmt"
	// "io/ioutil"
	// "testing"
)

type ArticleCategoryRepository interface {
	SaveArticleCategory(articleID uint,categoryID uint,categoryName string) (articleCategory models.ArticleCategory)
	GetArticleCategoryList(articleID uint) (articleCategories []models.ArticleCategory)
}

func NewArticleCategoryRepository() ArticleCategoryRepository {
	return &articleCategoryRepository{}
}

type articleCategoryRepository struct{}

func (n articleCategoryRepository)SaveArticleCategory(articleID uint,categoryID uint,categoryName string)(articleCategory models.ArticleCategory){
	articleCategory.ArticleID = articleID
	articleCategory.CategoryID = categoryID
	articleCategory.CategoryName = categoryName
	db := datasource.GetDB()
	db.Save(&articleCategory)
	return
}
func (n articleCategoryRepository)GetArticleCategoryList(articleID uint)(articleCategories []models.ArticleCategory){
	db := datasource.GetDB()
	db.Where("article_id = ?", articleID).Find(&articleCategories)
	return
}
