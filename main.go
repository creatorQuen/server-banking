package main

import (
	"ashishi-banking/app"
	"ashishi-banking/logger"
)

func main() {
	//log.Println("Starting our application... ")
	logger.Info("Starting our application... ")
	app.Start()
}
