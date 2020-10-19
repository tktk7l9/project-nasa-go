package main

import (
	"github.com/labstack/echo"
	"github.com/tktk7l9/project-nasa-go/routes"
)

func main() {
	e := echo.New()

	// Routes
	routes.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
