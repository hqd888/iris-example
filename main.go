package main

import (
	"github.com/hqd888/iris-example/web/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	template := iris.HTML("./web/views", ".html").Layout("layout/layout.html").Reload(true)
	app.RegisterView(template)
	app.StaticWeb("/public", "./web/public")

	mvc.New(app.Party("/demo")).Handle(new(controllers.DemoController))

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问页面出错！"))
		ctx.ViewLayout("")
		ctx.View("layout/error.html")
	})
	app.Run(iris.Addr("localhost:8080"))
}
