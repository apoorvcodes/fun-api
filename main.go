package main

import (

	"os"

	"github.com/TeamIndex/Backend/router"
	"github.com/gominima/cors"
	middleware "github.com/gominima/middlewares"
	"github.com/gominima/minima"

)

func main() {

	app := minima.New()
	app.Use(cors.New().AllowAll())
	app.UseRaw(middleware.Logger)
	app.Static("/static", "./Static")
	app.NotFound(func(res *minima.Response, req *minima.Request) {
		res.Render("./Static/notfound.html", req.Path())
	})
	app.UseRouter(router.MainRouter())
	app.Get("/", func(res *minima.Response, req *minima.Request) {
		res.OK().Send("Hello World")
	})
	app.Listen(":"+GetENV("PORT"))
}


func GetENV(key string) string {
	return os.Getenv(key)
}
