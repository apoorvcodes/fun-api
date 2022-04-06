package main

import (
	"log"
	"os"

	"github.com/TeamIndex/Backend/router"
	"github.com/gominima/cors"
	middleware "github.com/gominima/middlewares"
	"github.com/gominima/minima"
	"github.com/joho/godotenv"
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

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetENV(key string) string {
	return os.Getenv(key)
}
