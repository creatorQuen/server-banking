package main

import (
	"ashishi-banking/app"
	"ashishi-banking/logger"
)

func main() {
	//log.Println("Starting our application... ")
	logger.Log.Info("Starting our application... ")
	app.Start()
}
