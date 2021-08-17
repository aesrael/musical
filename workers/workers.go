package workers

import (
	"context"
	"fmt"
	"musical/config"

	"github.com/apex/log"
	"github.com/hibiken/asynq"
)

var JWT_KEY = config.Config["JWT_KEY"]
var REDIS_HOST = config.Config["REDIS_HOST"]
var REDIS_USERNAME = config.Config["REDIS_USERNAME"]
var REDIS_PORT = config.Config["REDIS_PORT"]
var REDIS_PASSWORD = config.Config["REDIS_PASSWORD"]

func main() {
	Addr := fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PASSWORD)

	conn := asynq.RedisClientOpt{
		Addr:     Addr,
		Password: REDIS_PASSWORD,
		Username: REDIS_USERNAME,
	}

	worker := asynq.NewServer(conn, asynq.Config{
		Concurrency: 10,
	})

	// Create a new task's mux instance.
	mux := asynq.NewServeMux()

	mux.HandleFunc(config.TASK_TYPE,
		func(c context.Context, t *asynq.Task) error {
			return nil
		})

	if err := worker.Run(mux); err != nil {
		log.Error(err.Error())
	}
}
