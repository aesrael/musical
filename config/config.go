package config

import (
	"os"

	"github.com/apex/log"
	"github.com/joho/godotenv"
)

type ConfigType map[string]string

var Config = ConfigType{
	"JWT_KEY":                "",
	"REDIS_HOST":             "",
	"REDIS_PORT":             "",
	"REDIS_PASSWORD":         "",
	"REDIS_USERNAME":         "musical",
	"GOOGLE_API_CREDENTIALS": "",
	"GOOGLE_DRIVE_FOLDER":    "",
}

const SERVER_PORT = ":8999"
const DL_TRACK_JOB = "job:download"
const UL_TRACK_JOB = "job:upload"
const BACKUP_DB_JOB = "job:backup"
const DB_FILE = "db.json"
const ALLOWED_ORIGINS = "https://github.com"

func InitConfig() {
	if err := godotenv.Load(); err != nil {
		log.Info("No .env file found")
	}

	env := os.Getenv("GO_ENV")
	log.Info("env: " + env)
	required := []string{
		"GO_ENV",
		"JWT_KEY",
		"REDIS_HOST",
		"REDIS_PORT",
		"REDIS_PASSWORD",
		"REDIS_USERNAME",
		"GITHUB_TOKEN",
		"GOOGLE_API_CREDENTIALS",
		"GOOGLE_DRIVE_FOLDER",
	}

	for _, env := range required {
		envVal, exists := os.LookupEnv(env)
		if !exists {
			log.Fatal(env + " not found in env")
		}
		if _, ok := Config[env]; ok {
			Config[env] = envVal
		}
	}
}
