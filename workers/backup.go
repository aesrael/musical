package workers

import (
	"context"
	"encoding/json"
	"musical/config"
	"musical/drive"

	"github.com/hibiken/asynq"
)

func backupDb(c context.Context, t *asynq.Task) error {
	keys, _ := RedisDB.Keys("*").Result()
	data := map[string]string{}

	for _, key := range keys {
		val, err := RedisDB.Get(key).Result()

		if err != nil && !isRedisNilError(err) && !isWrongTypeError(err) {
			return err
		}
		// empty or non string type, skip
		if val == "" {
			continue
		}
		data[key] = val
	}

	JSON, err := json.Marshal(data)

	if err != nil {
		return err
	}

	dbFileId, err := RedisDB.Get(config.DB_FILE).Result()

	if err != nil && !isRedisNilError(err) {
		return err
	}

	if dbFileId != "" {
		if err := drive.UpdateDriveFile(JSON, dbFileId); err != nil {
			return err
		}
		return nil
	}

	id, err := drive.UploadFileToDrive(JSON, config.DB_FILE)
	if err != nil {
		return err
	}

	if err := RedisDB.Set(config.DB_FILE, id, 0).Err(); err != nil {
		return err
	}

	return nil
}
