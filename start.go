package main

import "musical/config"

func initWebServer() {
	server := httpServer()
	server.Listen(config.SERVER_PORT)
}
