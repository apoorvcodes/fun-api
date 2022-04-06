package router

import (
	"github.com/TeamIndex/Backend/utils"
	"github.com/gominima/minima"
	"math/rand"
)

func JokeRoute() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		Type := req.Param("type")
		switch Type {
		case "all":
			jokes := utils.ParseJokes()
			res.OK().JSON(&map[string]interface{}{
				"Status": 200,
				"Jokes":  jokes,
			})
			res.CloseConn()
		case "single":
			jokes := utils.ParseJokes()
			integer := rand.Intn(len(jokes))
			res.OK().JSON(&map[string]interface{}{
				"Status": 200,
				"Jokes":  jokes[integer],
			})
			res.CloseConn()
		default:
			res.NotFound().JSON(&map[string]interface{}{
				"Status": 200,
				"Error":  "No params match All/Single",
			})
			res.CloseConn()
		}
	}
}
