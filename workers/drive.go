package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"musical/config"
	"net/http"
	"os"

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

func uploadFileToDrive(file *os.File, trackName string) error {
	client := driveClient()

	f := &drive.File{
		Name:    trackName,
		Parents: []string{config.Config["GOOGLE_DRIVE_FOLDER"]},
	}

	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return err
	}

	res, err := srv.Files.
		Create(f).
		Media(file).
		ProgressUpdater(func(now, size int64) { log.Info(fmt.Sprintf("%d, %d\r", now, size)) }).
		Do()

	if err != nil {
		return err
	}
	log.Info(fmt.Sprintf("file (%s) uploaded", res.Id))
	return nil
}
