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
	SaveArticleCategory(articleID uint,categoryID uint) (articleCategory models.ArticleCategory)
}

func NewArticleCategoryRepository() ArticleCategoryRepository {
	return &articleCategoryRepository{}
}

type articleCategoryRepository struct{}

func (n articleCategoryRepository)SaveArticleCategory(articleID uint,categoryID uint)(articleCategory models.ArticleCategory){
	articleCategory.ArticleID = articleID
	articleCategory.CategoryID = categoryID
	db := datasource.GetDB()
	db.Save(&articleCategory)
	return
}
