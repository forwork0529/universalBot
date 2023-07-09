package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type envVars struct{
	DebugLevel string `envconfig:"DEBUG_LEVEL" required:"true"`
}

var EnvVars envVars


func LoadConfig(path string){
	// читаем файл .env и присваиваем считанные переменные окружения процессу
	_ = godotenv.Load(path)
	// заполняем структуру переменными окружения процесса
	err := envconfig.Process("", &EnvVars)
	if err != nil{
		log.Fatalf("envconfig.Process(): %v\n", err.Error())
	}
}