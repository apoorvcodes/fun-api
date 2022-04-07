package router

import (
	// "fmt"

	"fmt"
	"os"

	"github.com/gominima/go-requests"
	"github.com/gominima/minima"
)

func NewsRoute() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		q := req.Param("query")
		data, err:= goquests.Get(goquests.Request{
			URL: fmt.Sprintf("https://newsapi.org/v2/everything?q=%v&apiKey=%v", q, os.Getenv("NEWS")),
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