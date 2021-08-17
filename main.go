package main

import (
	"musical/config"
)

func main() {
	config.InitConfig()
	InitQueue()
	initWebServer()
}
