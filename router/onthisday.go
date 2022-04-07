package router

import (
	"fmt"
	"time"

	goquests "github.com/gominima/go-requests"
	"github.com/gominima/minima"
	"github.com/TeamIndex/Backend/utils"
)


func OnThisDayRoute() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		typ := req.Param("type")
		_, month, day := time.Now().Date()
		data, err:= goquests.Get(goquests.Request{
			URL: fmt.Sprintf("https://en.wikipedia.org/api/rest_v1/feed/onthisday/%v/%v/%v", typ, utils.GetMonth(month.String()), day),
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