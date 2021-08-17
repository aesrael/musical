package main

import (
	"fmt"
	"musical/config"

	"github.com/hibiken/asynq"
)

var JWT_KEY = config.Config["JWT_KEY"]
var REDIS_HOST = config.Config["REDIS_HOST"]
var REDIS_USERNAME = config.Config["REDIS_USERNAME"]
var REDIS_PORT = config.Config["REDIS_PORT"]
var REDIS_PASSWORD = config.Config["REDIS_PASSWORD"]

var QClient *asynq.Client

func InitQueue() {
	Addr := fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PASSWORD)

	conn := asynq.RedisClientOpt{
		Addr:     Addr,
		Password: REDIS_PASSWORD,
		Username: REDIS_USERNAME,
	}

	QClient = asynq.NewClient(conn)
}
