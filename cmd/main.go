package main

import "botApi/pkg/logger"

func main(){
	logger.New()
	logger.Info("logger initialized")
}