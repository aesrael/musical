package workers

import (
	"fmt"
	"musical/config"

	"github.com/apex/log"
	"github.com/hibiken/asynq"
)

func InitWorkers() {
	REDIS_HOST := config.Config["REDIS_HOST"]
	REDIS_USERNAME := config.Config["REDIS_USERNAME"]
	REDIS_PORT := config.Config["REDIS_PORT"]
	REDIS_PASSWORD := config.Config["REDIS_PASSWORD"]

	Addr := fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PORT)
	fmt.Print(Addr)
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

	mux.HandleFunc(config.TASK_TYPE, processJob)

	if err := worker.Run(mux); err != nil {
		log.Error(err.Error())
	}
}
