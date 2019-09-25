package repo

import (
	"../datasource"
	"../models"
	// "../util"
	"github.com/spf13/cast"
	// "log"
	// "fmt"
	// "io/ioutil"
	// "testing"
)

type CategoryRepository interface {
	SaveCategory(m map[string]interface{})(category models.Category)
	GetCategoryList(m map[string]interface{})(categorys []models.Category)
	GetCategory(categoryID uint)(category models.Category)
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

type categoryRepository struct{}

func (n categoryRepository)SaveCategory(m map[string]interface{})(category models.Category){
	category.Name = cast.ToString(m["Name"]);
	db := datasource.GetDB()
	db.Save(&category);
	return
}

func (n categoryRepository)GetCategory(categoryID uint)(category models.Category){
	db := datasource.GetDB()
	db.First(&category, categoryID)
	return
}

func (n categoryRepository)GetCategoryList(m map[string]interface{})(categorys []models.Category){
	db := datasource.GetDB()
	err := db.Find(&categorys).Error
	if err!= nil{
		panic("select Error")
	}
	return
}
