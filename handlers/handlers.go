package handlers

import (
	"log"

	"github.com/centrifugal/centrifuge"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	node *centrifuge.Node
}

func NewHandler(node *centrifuge.Node) *Handler {
	return &Handler{node: node}
}

func (h *Handler) WebsocketHandler(c echo.Context) error {
	h.HandleClientConnections()

	wsHandler := centrifuge.NewWebsocketHandler(h.node, centrifuge.WebsocketConfig{})

	wsHandler.ServeHTTP(c.Response(), c.Request())

	return nil
}

func (h *Handler) HandleClientConnections() {
	h.node.OnConnect(func(client *centrifuge.Client) {
		transportName := client.Transport().Name()
		transportProto := client.Transport().Protocol()
		log.Printf("client connected via %s (%s)", transportName, transportProto)

		client.OnSubscribe(func(e centrifuge.SubscribeEvent, cb centrifuge.SubscribeCallback) {
			log.Printf("client subscribes on channel %s", e.Channel)
			cb(centrifuge.SubscribeReply{}, nil)
		})

		client.OnPublish(func(e centrifuge.PublishEvent, cb centrifuge.PublishCallback) {
			log.Printf("client publishes into channel %s: %s", e.Channel, string(e.Data))
			cb(centrifuge.PublishReply{}, nil)
		})

		client.OnDisconnect(func(e centrifuge.DisconnectEvent) {
			log.Printf("client disconnected")
		})
	})
}
