package router

import (
	goquests "github.com/gominima/go-requests"
	"github.com/gominima/minima"
)

func FactRoute() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		data, err:= goquests.Get(goquests.Request{
			URL: "https://www.thefact.space/random",
			Headers: make(map[string]string),
			Data: make(map[string]interface{}),
		})
		if err != nil {
			res.NotFound().JSON(map[string]interface{}{
				"Status": 404,
				"Error": err.Error,
			})
			return
		}
		res.OK().JSON(map[string]interface{}{
			"Status": 200,
			"Data": data.Body,
		})
	}
}