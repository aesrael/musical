package workers

import (
	"fmt"
	"musical/config"
	"musical/server"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/apex/log"
	"github.com/hibiken/asynq"
)

const SPOTDL = "spotdl"

const SPOTIFY_TRACK_REGEX = `https://open.spotify.com/(track|playlist|album|artist)/\w*`
const ROOT_FILES_PATH = "files/"

func downloadTrack(track string) error {
	// only spotify tracks are allowed for now.
	matched, err := regexp.MatchString(SPOTIFY_TRACK_REGEX, track)
	if !matched {
		log.Warn("Invalid track url: " + track)
		return err
	}

	log.WithField("track", track).Info("downloading track")

	cmd := exec.Command(SPOTDL, track)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	trackID, err := GetTrackId(track)
	if err != nil {
		return err
	}

	dir := filepath.Join(ROOT_FILES_PATH, trackID)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	cmd.Dir = dir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("exec error, %w", asynq.SkipRetry)
	}

	log.WithField("track", track).Info("track succesfully downloaded")

	if err := initPostDownloadActions(trackID); err != nil {
		return err
	}
	return nil
}

// prep track for upload to cloud
func initPostDownloadActions(trackID string) error {
	log.WithField("trackId", trackID).Info("enqueing new upload job")
	// push record to db. key is track ID(dir name), value is status
	// +------------------------+---------+
	// |          key           |  value  |
	// +------------------------+---------+
	// | 0VjIjW4GlUZAMYd2vXMi3b | pending |
	// +------------------------+---------+
	if err := RedisDB.Set(trackID, "pending", 0).Err(); err != nil {
		return err
	}

	// add task to upload queue
	job := asynq.NewTask(config.UL_TRACK_JOB, []byte(trackID))
	if _, err := server.QClient.Enqueue(job); err != nil {
		return err
	}

	return nil
}
