package main

import (
	"fmt"
	"musical/config"

	"github.com/hibiken/asynq"
)

var QClient *asynq.Client

func InitQueue() {
	REDIS_HOST := config.Config["REDIS_HOST"]
	REDIS_USERNAME := config.Config["REDIS_USERNAME"]
	REDIS_PORT := config.Config["REDIS_PORT"]
	REDIS_PASSWORD := config.Config["REDIS_PASSWORD"]

	Addr := fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PORT)

	conn := asynq.RedisClientOpt{
		Addr:     Addr,
		Password: REDIS_PASSWORD,
		Username: REDIS_USERNAME,
	}

	QClient = asynq.NewClient(conn)
}
