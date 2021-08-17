package main

import (
	"musical/config"
	"musical/workers"
	"sync"
)

func main() {
	config.InitConfig()
	InitQueue()

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		initWebServer()
		wg.Done()
	}()

	go func() {
		workers.InitWorkers()
		wg.Done()
	}()

	wg.Wait()
}
