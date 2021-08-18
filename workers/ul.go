package workers

import (
	"os"
	"path/filepath"

	"github.com/apex/log"
)

func uploadTrack(dirName string) error {
	dirPath := filepath.Join(ROOT_FILES_PATH, dirName)
	f, err := os.Open(dirPath)

	if err != nil {
		return err
	}

	files, err := f.Readdir(0)
	if err != nil {
		return err
	}

	for _, v := range files {
		if !isMusicFile(v.Name()) {
			continue
		}
		trackPath := filepath.Join(dirPath, v.Name())
		file, err := os.Open(trackPath)
		if err != nil {
			return err
		}

		if err := uploadFileToDrive(file, v.Name()); err != nil {
			return err
		}
		log.Info("file uploded: " + v.Name())
	}

	// upload track to drive
	if err := RedisDB.Set(dirName, "done", 0).Err(); err != nil {
		return err
	}

	return nil
}
