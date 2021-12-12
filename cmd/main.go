package main

import (
	"anderson.co/web_service/shared/server"
	"anderson.co/web_service/shared/server/fiber"
)

type app struct {
	webServer server.WebServer
}

func NewApp(webServer server.WebServer) *app {
	return &app{ webServer }
}

func main() {
	fiberWebServer := fiber.NewServer()
	app := NewApp(fiberWebServer)
	app.webServer.Run("3030")
}
