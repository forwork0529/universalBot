package main

import (
	"botApi/pkg/config"
	"botApi/pkg/logger"
)

func main(){

	logger.New()
	logger.Info("logger initialized")

	config.LoadConfig("/.env")
	logger.Info("environment variables initialized")

	shutDownChannel := make(chan struct{})


}