package main

import (
	"musical/config"
	"musical/server"
	"musical/workers"
	"sync"
)

func main() {
	config.InitConfig()
	server.InitQueue()

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		server.InitWebServer()
		wg.Done()
	}()

	go func() {
		workers.InitWorkers()
		wg.Done()
	}()

	wg.Wait()
}
