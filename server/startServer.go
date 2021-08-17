package server

import "musical/config"

func InitWebServer() {
	server := httpServer()
	server.Listen(config.SERVER_PORT)
}
