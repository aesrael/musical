package workers

import (
	"context"
	"encoding/json"
	"musical/server"

	"github.com/apex/log"
	"github.com/hibiken/asynq"
)

func processJob(c context.Context, t *asynq.Task) error {
	params := &server.JobParams{}

	jobByte := t.Payload()
	err := json.Unmarshal(jobByte, params)
	if err != nil {
		return nil
	}
	log.Info("processing job")
	return nil
}
