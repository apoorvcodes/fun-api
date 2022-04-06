package main

import (
	"github.com/gominima/cors"
	middleware "github.com/gominima/middlewares"
	"github.com/gominima/minima"
	"github.com/TeamIndex/Backend/router"
)

func main() {
 Init()
 app := minima.New()
 app.Use(cors.New().Default())
 app.UseRaw(middleware.Logger)
 app.UseRouter(router.MainRouter())
 app.Get("/", func(res *minima.Response, req *minima.Request) {
	 res.OK().Send("Hello World")
 })
 app.Listen(GetENV("PORT"))
}