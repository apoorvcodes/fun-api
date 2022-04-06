package router

import "github.com/gominima/minima"

func JokeRoute() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		res.OK().Send("Hello World")
	}
}