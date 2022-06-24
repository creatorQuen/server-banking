package main

import (
	"ashishi-banking/app"
	"ashishi-banking/logger"
)

func main() {
	logger.Info("Starting our application... ")
	app.Start()
}
