package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"musical/server"

	"github.com/apex/log"
	"github.com/hibiken/asynq"
)

func processDownload(c context.Context, t *asynq.Task) error {
	job := &server.Job{}

	jobByte := t.Payload()
	if err := json.Unmarshal(jobByte, job); err != nil {
		return nil
	}

	trackId, err := GetTrackId(job.Track)
	if err != nil {
		return err
	}
	existing, err := RedisDB.Get(trackId).Result()

	if err != nil && !isRedisNilError(err) {
		return err
	}

	if existing == "done" || existing == "pending" {
		return fmt.Errorf("file already processed, %w", asynq.SkipRetry)
	}

	log.WithField("job", job).Info("processing job")

	if err := downloadTrack(job.Track); err != nil {
		log.WithField("track", job.Track).Error(fmt.Sprintf("download error: %s", err.Error()))
		return err
	}
	return nil
}

func processUpload(c context.Context, t *asynq.Task) error {
	job := string(t.Payload())

	log.WithField("job", job).Info("processing job " + job)

	if err := uploadTrack(job); err != nil {
		log.Error(fmt.Sprintf("upload error: %s", err.Error()))
		return err
	}
	return nil
}
