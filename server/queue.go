package server

import (
	"encoding/json"
	"fmt"
	"log"
	"musical/config"
	"musical/drive"

	"github.com/go-redis/redis/v7"
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

	RedisDB := conn.MakeRedisClient().(*redis.Client)

	if err := seedDB(RedisDB); err != nil {
		log.Fatal(err.Error())
	}

	QClient = asynq.NewClient(conn)
}

func seedDB(db *redis.Client) error {
	keys, _ := db.Keys("*").Result()
	if len(keys) != 0 {
		// return nil
	}

	res, err := drive.GetCloudDB()
	if err != nil {
		return err
	}

	if res == nil {
		return nil
	}

	defer res.Body.Close()

	var entries = &map[string]string{}

	jsonErr := json.NewDecoder(res.Body).Decode(entries)

	if jsonErr != nil {
		return err
	}

	for k, v := range *entries {
		if err := db.Set(k, v, 0).Err(); err != nil {
			return err
		}
	}

	return nil
}
