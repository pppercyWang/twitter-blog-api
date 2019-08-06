package models

import (
	"github.com/jinzhu/gorm"
	"time"
)
type Wechat struct {
	gorm.Model
	//AgencyID             uint		//代理后台用户ID
	CustomerID             uint		//客户端用户ID
	WechatID           string `gorm:"not null;unique" sql:"index"`
	WechatPassword     string
	SocksIp            string
	WechatAlias        string
	WechatMobile       string
	Nickname           string
	Sex                uint
	Signature          string
	Country            string
	Province           string
	City               string
	SmallHeadimgUrl    string
	GroupID            uint
	LastLoginDate      time.Time
	LoginDeviceData    string
	Remark             string
	Status             int
}

