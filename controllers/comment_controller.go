package controllers

import (
	"github.com/kataras/iris"
	"log"
	"../models"
	"../service"
	// "fmt"
)
type CommentController struct {
	Ctx     iris.Context
	Service service.CommentService
}
func NewCommentController() *CommentController {
	return &CommentController{ Service: service.NewCommentService() }
}
func (g *CommentController) PostSave() (result models.Result)  {
	var m map[string]interface{}
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	if m["Content"] == "" || m["Content"] == nil {
		result.Code = -1
		result.Msg = "评论内容不能为空"
		return
	}
	if m["ArticleID"] == "" || m["ArticleID"] == nil {
		result.Code = -1
		result.Msg = "文章ID不能为空"
		return
	}
	if m["GitUserID"] == "" || m["GitUserID"] == nil {
		result.Code = -1
		result.Msg = "用户ID不能为空"
		return
	}
	if m["Username"] == "" || m["Username"] == nil {
		result.Code = -1
		result.Msg = "用户名不能为空"
		return
	}
	if m["AvatarUrl"] == "" || m["AvatarUrl"] == nil {
		result.Code = -1
		result.Msg = "头像链接不能为空"
		return
	}
	if m["GithubUrl"] == "" || m["GithubUrl"] == nil {
		result.Code = -1
		result.Msg = "github链接不能为空"
		return
	}
	return g.Service.SaveComment(m)
}
func (g *CommentController) PostList() (result models.Result)  {
	var m map[string]interface{}
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	return g.Service.GetCommentList(m)
}

