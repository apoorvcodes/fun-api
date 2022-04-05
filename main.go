package main

import (
	"github.com/gominima/cors"
	middleware "github.com/gominima/middlewares"
	"github.com/gominima/minima"
)

func main() {
 app := minima.New()
 app.Use(cors.New().Default())
 app.UseRaw(middleware.Logger)
 app.Get("/", func(res *minima.Response, req *minima.Request) {
	 res.OK().Send("Hello World")
 })
 app.Listen(":3000")
}