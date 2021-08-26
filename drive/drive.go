package drive

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"musical/config"
	"net/http"

	"github.com/apex/log"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func driveClient() *http.Client {
	jsonBytes := []byte(config.Config["GOOGLE_API_CREDENTIALS"])

	var c = struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}{}
	json.Unmarshal(jsonBytes, &c)
	config := &jwt.Config{
		Email:      c.Email,
		PrivateKey: []byte(c.PrivateKey),
		Scopes: []string{
			drive.DriveScope,
		},
		TokenURL: google.JWTTokenURL,
	}
	client := config.Client(context.Background())
	return client
}

func UploadFileToDrive(file []byte, fileName string) (string, error) {
	client := driveClient()

	f := &drive.File{
		Name:    fileName,
		Parents: []string{config.Config["GOOGLE_DRIVE_FOLDER"]},
	}

	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return "", err
	}

	res, err := srv.Files.
		Create(f).
		Media(bytes.NewReader(file)).
		ProgressUpdater(func(now, size int64) { log.Info(fmt.Sprintf("%d, %d\r", now, size)) }).
		Do()

	if err != nil {
		return "", err
	}
	log.Info(fmt.Sprintf("file (%s) uploaded", res.Id))
	return res.Id, nil
}

func UpdateDriveFile(file []byte, fileId string) error {
	client := driveClient()

	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return err
	}

	res, err := srv.Files.
		Update(fileId, new(drive.File)).
		Media(bytes.NewReader(file)).
		ProgressUpdater(func(now, size int64) { log.Info(fmt.Sprintf("%d, %d\r", now, size)) }).
		Do()

	if err != nil {
		return err
	}
	log.Info(fmt.Sprintf("file (%s) updated", res.Id))
	return nil
}

func GetCloudDB() (*http.Response, error) {
	client := driveClient()

	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	q := fmt.Sprintf("name='db.json' and parents in '%s'", config.Config["GOOGLE_DRIVE_FOLDER"])
	res, err := srv.Files.
		List().
		Q(q).
		Fields("nextPageToken, files(id, name)").Do()

	if err != nil {
		return nil, err
	}

	if len(res.Files) == 0 {
		return nil, nil
	}

	if len(res.Files) > 1 {
		return nil, errors.New("conflicting db files")
	}

	resp, err := srv.Files.Get(res.Files[0].Id).Download()
	return resp, err
}
