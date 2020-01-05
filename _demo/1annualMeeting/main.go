package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

//type lotteryController struct {
//	Ctx iris.Context
//}

func main() {
	app := iris.Default()

	html := iris.HTML("/Users/phoenix/WorkspaceGo/lottery/_demo/1annualMeeting", ".html")
	app.RegisterView(html)

	app.Get("/", func(context iris.Context) {
		context.WriteString("Hello World! from iris")
	})
	app.Get("hello", func(context context.Context) {
		context.ViewData("Title", "Test page!")
		context.ViewData("Content", "aa")
		context.View("hello.html")
	})
	println("c")
	app.Run(iris.Addr(":8088"), iris.WithCharset("utf-8"))
}
