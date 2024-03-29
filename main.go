package main

import (
	"sync"

	"github.com/mesxx/Echo_Simple_Self_Payroll_API/config"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Infof(".env is not loaded properly")
	}
	log.Infof("read .env from file")

	config := config.NewConfig()
	server := InitServer(config)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Run()
	}()

	wg.Wait()
}
