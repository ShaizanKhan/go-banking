package main

import (
	"github.com/ShaizanKhan/go-banking-lib/logger"
	"github.com/ShaizanKhan/go-banking/app"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
