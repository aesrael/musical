package server

import "musical/config"

func InitWebServer() {
	defer QClient.Close()
	server := httpServer()
	server.Listen(config.SERVER_PORT)
}
