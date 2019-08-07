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
	app.Use(middleware.GetJWT().Serve)
	mvc.New(app.Party( bathUrl + "/article")).Handle(controllers.NewArticleController())
}
