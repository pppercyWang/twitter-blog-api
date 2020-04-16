package service

import (
	"../models"
	"../repo"
	// "github.com/russross/blackfriday"
)
type CommentService interface {
	SaveComment(m map[string]interface{}) (result models.Result)
	GetCommentList(m map[string]interface{}) (result models.Result)
}
type commentService struct {
}
func NewCommentService() CommentService {
	return &commentService{}
}
var commentRepo = repo.NewCommentRepository()
func (u commentService) SaveComment(m map[string]interface{}) (result models.Result){
	result.Code = 0
	comment := commentRepo.SaveComment(m)
	maps := make(map[string]interface{},1)
	maps["comment"] = comment
	result.Data = maps
	result.Msg = "评论成功"
	return
}
func (u commentService) GetCommentList(m map[string]interface{}) (result models.Result){
	total,comments := commentRepo.GetCommentList(m)
	result.Code = 0
	maps := make(map[string]interface{},1)
	if total != 0 {
		maps["Total"] = total
	}
	maps["List"] = comments
	result.Data = maps
	return
}