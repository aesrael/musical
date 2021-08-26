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

	trackId, err := GetTrackId(job.Title)
	if err != nil {
		return err
	}
	res, err := RedisDB.Get(trackId).Result()

	if err != nil && !isRedisNilError(err) {
		return err
	}

	if res != "" {
		return fmt.Errorf("file already processed, %w", asynq.SkipRetry)
	}

	log.WithField("job", job).Info("processing job #" + job.IssueNumber)

	if err := downloadTrack(job.Title); err != nil {
		log.WithField("track", job.Title).Error(fmt.Sprintf("download error: %s", err.Error()))
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
