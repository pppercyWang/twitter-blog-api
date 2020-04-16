package repo

import (
	"../datasource"
	"../models"
	// "../util"
	"github.com/spf13/cast"
	// "log"
	// "io/ioutil"
	// "testing"
)

type CommentRepository interface {
	SaveComment(m map[string]interface{})(comment models.Comment)
	GetCommentList(m map[string]interface{})(total int,comments []models.Comment)
	GetComment(commentID uint)(comment models.Comment)
}

func NewCommentRepository() CommentRepository {
	return &commentRepository{}
}
var articleRepo = NewArticleRepository()

type commentRepository struct{}

func (n commentRepository)SaveComment(m map[string]interface{})(comment models.Comment){
	db := datasource.GetDB()
	article := articleRepo.GetArticle(cast.ToUint(m["ArticleID"]) )
	comment.ArticleTitle = article.Title
	comment.Content = cast.ToString(m["Content"])
	comment.ArticleID =cast.ToUint(m["ArticleID"]) 
	comment.GitUserID = cast.ToUint(m["GitUserID"])
	comment.Username = cast.ToString(m["Username"])
	comment.AvatarUrl =cast.ToString( m["AvatarUrl"])
	comment.GithubUrl = cast.ToString(m["GithubUrl"])	
	db.Save(&comment);
	return
}

func (n commentRepository)GetComment(commentID uint)(comment models.Comment){
	db := datasource.GetDB()
	db.First(&comment, commentID)
	return
}

func (n commentRepository)GetCommentList(m map[string]interface{})(total int,comments []models.Comment){
	db := datasource.GetDB()
	var err error
	if m["ArticleID"] == nil {
		if m["Size"]==nil || m["Page"]==nil {
			err = db.Order("created_at desc").Find(&comments).Error
		}else{
			db.Model(&models.Comment{}).Count(&total)
			err = db.Limit(cast.ToInt(m["Size"])).Offset((cast.ToInt(m["Page"])-1)*cast.ToInt(m["Size"])).Order("created_at desc").Find(&comments).Error
		}
	}else{
		err = db.Where("article_id = ?", cast.ToUint(m["ArticleID"]) ).Order("created_at desc").Find(&comments).Error
	}

	if err!= nil{
		panic("select Error")
	}
	return
}
