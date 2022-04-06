package router

import "github.com/gominima/minima"

func MainRouter() *minima.Router {
 router := minima.NewRouter()
 router.Get("/joke", JokeRoute())
 return router
}