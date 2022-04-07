package router

import (
	"fmt"

	goquests "github.com/gominima/go-requests"
	"github.com/gominima/minima"
)



func MemeRouter() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		sub := req.Param("subreddit")
		count := req.Param("count")
        switch sub {
		case "all" :
			data, err:= goquests.Get(goquests.Request{
				URL: fmt.Sprintf("https://meme-api.herokuapp.com/gimme/%v", count),
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
		default:
			data, err:= goquests.Get(goquests.Request{
				URL: fmt.Sprintf("https://meme-api.herokuapp.com/gimme/%v/%v", sub, count),
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
}