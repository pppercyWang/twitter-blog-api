/*
@Time : 2019/4/8 10:03
@Author : lukebryan
@File : wechat_repo
@Software: GoLand
*/
package repo

import (
	"../datasource"
	"../models"
	"github.com/spf13/cast"
)

type WechatRepository interface {
	List(m map[string]interface{}) (total int, wechats []models.Wechat)
}

func NewWechatRepository() WechatRepository {
	return &wechatRepository{}
}

type wechatRepository struct{}

//条件分页查询
func (n wechatRepository) List(m map[string]interface{}) (total int, wechats []models.Wechat) {
	db := datasource.GetDB()
	err := db.Limit(cast.ToInt(m["Size"])).Offset((cast.ToInt(m["Page"])-1)*cast.ToInt(m["Size"])).Find(&wechats).Error
	if err!=nil {
		panic("select Error")
	}
	datasource.GetDB().Model(&models.Wechat{}).Count(&total)
	return
}

