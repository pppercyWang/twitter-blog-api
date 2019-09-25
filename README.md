### iris+gorm为twitter-blog搭建的服务
尽量把功能做的简单易懂,希望能给也想写个人博客的朋友一个参考


### 功能概述
1. 文章保存
2. 文章分页查询
3. 实现jwt登录接口 //用于提供写文章的权限
4. 保存文章和用户的事务处理


### 表结构
暂时只用到4张表

```
//文章表
type Article struct {
	gorm.Model
	Content        string `gorm:"type:text"` // 内容
	Title     string `gorm:"type:varchar(80);not null;"`  // 标题
	Tags  string `gorm:"type:varchar(20)"` // 文章所带标签
	Personal uint `gorm:"default:0"` // 是否是原创
	Categories string `gorm:"type:varchar(20)"` // 分类字符串 eg: node.js,vue,golang,react
}
```
```
// 文章分类表
type Category struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null;"`
}
```
```
// 多对多联接表
type ArticleCategory struct {
	gorm.Model
	CategoryID uint `gorm:"not null"`
	ArticleID uint `gorm:"not null"`
	CategoryName string `gorm:"type:varchar(20)"`
}
```
```
// 用户表
type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	Password     string
	Name         string                      
}
```