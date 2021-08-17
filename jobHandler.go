package main

import (
	"encoding/json"
	"musical/config"

	"net/http"

	"github.com/apex/log"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
)

type JobParams struct {
	IssueNumber string `json:"issue_number"`
	Title       string `json:"title"`
}

func enqueueJob(c *fiber.Ctx) error {
	params := &JobParams{}

	reqBody := c.Body()
	err := json.Unmarshal(reqBody, params)
	if err != nil {
		c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}

	log.WithField("job", params).Info("enqueing new job")
	// push job to queue from where it is eventually picked up and
	// processed by the worker
	job := asynq.NewTask(config.TASK_TYPE, reqBody)

	if _, err := QClient.Enqueue(job); err != nil {
		log.Error(err.Error())
	}
	log.WithField("job", params).Info("job queued succesfully")
	return nil
}