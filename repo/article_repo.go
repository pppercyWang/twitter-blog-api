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
)

type AticleRepository interface {
	GetArticleList(m map[string]interface{}) (total int, articles []models.Article)
}

func NewArticleRepository() AticleRepository {
	return &acticleRepository{}
}

type acticleRepository struct{}

func (n acticleRepository)GetArticleList(m map[string]interface{})(total int,articles []models.Article){
	db := datasource.GetDB()
	err := db.Limit(cast.ToInt(m["Size"])).Offset((cast.ToInt(m["Page"])-1)*cast.ToInt(m["Size"])).Find(&articles).Error
	if err!=nil {
		panic("select Error")
	}
	datasource.GetDB().Model(&models.Article{}).Count(&total)
	return
}

