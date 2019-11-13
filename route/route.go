package route

import (
	"../controllers"
	"../middleware"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func InitRouter(app *iris.Application) {
	bathUrl := "/api"
	mvc.New(app.Party(bathUrl + "/user")).Handle(controllers.NewUserController())
	mvc.New(app.Party( bathUrl + "/category")).Handle(controllers.NewCategoryController())
	mvc.New(app.Party( bathUrl + "/article")).Handle(controllers.NewArticleController())
	mvc.New(app.Party( bathUrl + "/auth")).Handle(controllers.NewAuthController())
	mvc.New(app.Party( bathUrl + "/comment")).Handle(controllers.NewCommentController())
	app.Use(middleware.GetJWT().Serve)
}
