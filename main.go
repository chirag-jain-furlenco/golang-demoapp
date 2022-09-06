package main

import (
	"demoapp/api"
	"demoapp/config"
	"demoapp/db"
	"sync"
)

func init() {
	var wg sync.WaitGroup
	wg.Add(1)

	config.InitializeConfig()
	go db.ConnectToDb(&wg)

	wg.Wait()
}

func main() {
	api.StartServer()
}
