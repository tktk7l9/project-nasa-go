package routes

import (
	"github.com/labstack/echo"
	"github.com/tktk7l9/project-nasa-go/web/api"
)

func Init(e *echo.Echo) {

	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVideos())
	}

}
