package main

import "musical/config"

func initWebServer() {
	server := httpServer()
	addr := "localhost" + config.SERVER_PORT
	server.Listen(addr)
}
