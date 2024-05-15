package main

import (
	"go.uber.org/zap"
)

func main() {

	//logger := zap.Must(zap.NewProduction())

	logger := zap.Must(zap.NewDevelopment())

	defer logger.Sync()

	logger.Info("Hello from Zap logger!")

}
