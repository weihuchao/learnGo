package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewProduction()
	logger.Info("Zap log info.")

}
