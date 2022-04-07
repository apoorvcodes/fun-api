package router

import (
	goquests "github.com/gominima/go-requests"
	"github.com/gominima/minima"
)

func ShortStoryRoute() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		q := req.Query("q")
		switch q {
		case "all":
			data, err:= goquests.Get(goquests.Request{
				URL: "https://shortstories-api.herokuapp.com/stories",
				Headers: make(map[string]string),
				Data: make(map[string]interface{}),
				
			})
			if err != nil {
				res.NotFound().JSON(map[string]interface{}{
					"Status": 404,
					"Error": err.Error,
				})
			}
			res.OK().JSON(map[string]interface{}{
				"Status": 200,
				"Data": data.Body,
			})
		default:
			data, err:= goquests.Get(goquests.Request{
				URL: "https://shortstories-api.herokuapp.com/",
				Headers: make(map[string]string),
				Data: make(map[string]interface{}),
				
			})
			if err != nil {
				res.NotFound().JSON(map[string]interface{}{
					"Status": 404,
					"Error": err.Error,
				})
			}
			res.OK().JSON(map[string]interface{}{
				"Status": 200,
				"Data": data.Body,
			})

		}
	}
}