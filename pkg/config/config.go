package config

import (
	"log"

	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

type Config struct {
	Dir  string `json:"dir"`
	Port string `json:"port"`
}

var config Config

func initConfig() map[string]string {
	var err error
	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Panic(err)
	}

	return envs
}

func GetConfig() Config {
	onceBody := func() {
		envs := initConfig()

		config = Config{
			Dir:  envs["DIR"],
			Port: envs["PORT"],
		}
	}

	once.Do(onceBody)

	return config
}
