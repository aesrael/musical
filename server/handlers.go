package server

import (
	"encoding/json"
	"musical/config"

	"net/http"

	"github.com/apex/log"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
)

type Job struct {
	Track string `json:"track"`
}

func enqueueJob(c *fiber.Ctx) error {
	params := &Job{}

	reqBody := c.Body()
	err := json.Unmarshal(reqBody, params)
	if err != nil {
		c.Status(http.StatusBadRequest).SendString(err.Error())
		log.Error(err.Error())
		return err
	}

	log.WithField("job", params).Info("enqueing new job")
	// push job to queue from where it is eventually picked up and
	// processed by the worker
	job := asynq.NewTask(config.DL_TRACK_JOB, reqBody)

	if _, err := QClient.Enqueue(job); err != nil {
		log.Error(err.Error())
		return err
	}
	log.WithField("job", params).Info("job queued succesfully")
	return c.SendStatus(http.StatusOK)
}

func backupDb(c *fiber.Ctx) error {
	log.Info("uploading new db file")
	job := asynq.NewTask(config.BACKUP_DB_JOB, c.Body())

	if _, err := QClient.Enqueue(job); err != nil {
		log.Error(err.Error())
		c.Status(http.StatusInternalServerError).SendString(err.Error())
		return err
	}
	return c.SendStatus(http.StatusOK)
}
