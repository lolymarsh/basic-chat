package main

import (
	"log"
	"loly/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.SetupWebSocket(e)

	e.Static("/", "./public")

	log.Printf("Starting server, visit http://localhost:8000")
	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
	}
}
