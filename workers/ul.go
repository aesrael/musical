package workers

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"musical/drive"
	"os"
	"path/filepath"
	"sync"

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

	wg := new(sync.WaitGroup)
	wg.Add(len(files))

	anyFailed := false

	for _, v := range files {
		go asyncUpload(v, wg, dirPath, &anyFailed)
	}

	wg.Wait()

	if anyFailed {
		return errors.New("upload error: some files not uploaded")
	}

	if err := RedisDB.Set(dirName, "done", 0).Err(); err != nil {
		return err
	}

	return os.RemoveAll(dirPath)
}

func asyncUpload(v fs.FileInfo, wg *sync.WaitGroup, dirPath string, anyFailed *bool) {
	defer wg.Done()

	if !isMusicFile(v.Name()) {
		return
	}

	trackPath := filepath.Join(dirPath, v.Name())
	file, err := ioutil.ReadFile(trackPath)

	if err != nil {
		log.Error(fmt.Sprintf("upload error occured %s", err.Error()))
		*anyFailed = true
		return
	}

	_, uploadErr := drive.UploadFileToDrive(file, v.Name())

	if uploadErr != nil {
		log.Error(fmt.Sprintf("upload error occured %s", err.Error()))
		*anyFailed = true
		return
	}

	log.Info("file uploded: " + v.Name())
	os.RemoveAll(trackPath)
}
