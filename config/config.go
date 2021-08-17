package config

import (
	"os"

	"github.com/apex/log"
	"github.com/joho/godotenv"
)

type ConfigType map[string]string

var Config = ConfigType{
	"JWT_KEY":        "",
	"REDIS_HOST":     "",
	"REDIS_PORT":     "",
	"REDIS_PASSWORD": "",
	"REDIS_USERNAME": "musical",
}

const SERVER_PORT = "8999"
const TASK_TYPE = "job:issue"

func InitConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	env, _ := os.LookupEnv("GO_ENV")
	log.Info("env: " + env)
	required := []string{
		"GO_ENV",
		"JWT_KEY",
	}

	for _, env := range required {
		envVar, exists := os.LookupEnv(env)
		if !exists {
			log.Fatal(env + " not found in env")
		}
		if val, ok := Config[envVar]; ok {
			Config[envVar] = val
		}
	}
}
