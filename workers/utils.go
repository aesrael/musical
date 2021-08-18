package workers

import (
	"errors"
	"net/url"
	"strings"
)

// split spotify, obtain track ID
func GetTrackId(trackURL string) (string, error) {
	URL, err := url.Parse(trackURL)
	if err != nil {
		return "", err
	}

	parts := strings.Split(URL.Path, "/")
	if len(parts) < 3 {
		return "", errors.New("Invalid URL: " + trackURL)
	}

	return parts[len(parts)-1], nil
}

var supportedAudioTypes = map[string]bool{
	"mp3": true,
	"M4A": true,
	"WAV": true,
	"WMA": true,
	"AAC": true,
}

func isMusicFile(trackTitle string) bool {
	parts := strings.Split(trackTitle, ".")
	ext := parts[len(parts)-1]

	return supportedAudioTypes[ext]
}
