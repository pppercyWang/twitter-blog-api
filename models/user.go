/*
@Time : 2019/5/8 10:05
@Author : lukebryan
@File : user
@Software: GoLand
*/
package models

import "github.com/jinzhu/gorm"

/*
用户端注册用户
 */
type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	Password     string
	Name         string                                         //姓名
	Email        string                         //邮箱
	Mobile       string `gorm:"unique"` //手机
	QQ           string
	Gender       int                         //0男 1女
	Age          int  //年龄
	Remark       string                      //备注
	Token      string `gorm:"-"`
	Session      string `gorm:"-"`
}
