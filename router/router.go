package router

import "github.com/gominima/minima"

func MainRouter() *minima.Router {
	router := minima.NewRouter()
	router.Get("/jokes/:type", JokeRoute())
	router.Get("/news/:query", NewsRoute())
    router.Get("/meme/:subreddit/:count", MemeRouter())
	router.Get("/onthisday/:type", OnThisDayRoute())
	router.Get("/story", ShortStoryRoute())
	router.Get("/facts", FactRoute())
	return router
}
