package routes

import (
	"log"
	"loly/handlers"
	"loly/middleware"

	"github.com/centrifugal/centrifuge"
	"github.com/labstack/echo/v4"
)

func SetupWebSocket(e *echo.Echo) {

	node, err := centrifuge.New(centrifuge.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := node.Run(); err != nil {
		log.Fatal(err)
	}

	h := handlers.NewHandler(node)

	e.GET("/connection/websocket", h.WebsocketHandler, middleware.AuthMiddleware)

}
