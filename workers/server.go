package workers

import (
	"fmt"
	"musical/config"
	"time"

	"github.com/apex/log"
	"github.com/go-redis/redis/v7"
	"github.com/hibiken/asynq"
)

var RedisDB *redis.Client

func InitWorkers() {
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

	RedisDB = conn.MakeRedisClient().(*redis.Client)

	worker := asynq.NewServer(conn, asynq.Config{
		Concurrency: 20,
		RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
			return 60 * time.Second
		},
	})

	mux := asynq.NewServeMux()

	mux.HandleFunc(config.DL_TRACK_JOB, processDownload)
	mux.HandleFunc(config.UL_TRACK_JOB, processUpload)
	mux.HandleFunc(config.BACKUP_DB_JOB, backupDb)

	if err := worker.Run(mux); err != nil {
		log.Error(err.Error())
	}
}
