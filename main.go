package main

import (
	"github.com/events-manager/application/routes"
	"github.com/events-manager/repository/factory"
	"github.com/events-manager/repository/rabbitmq"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.AddTrailingSlash())

	eventsRepository := rabbitmq.NewEvents(
		factory.GetRabbitConnection(),
	)

	routes.NewEvents(e, eventsRepository).Handler()
	// Start server
	e.Logger.Fatal(e.Start(":80"))
}
