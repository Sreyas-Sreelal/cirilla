package youtubedl

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os/exec"
)

//YtOptions holds options to be passed to youtubedl
type YtOptions struct {
	Path        string
	TrackName   string
	AudioFormat string
}

//TrackInfo holds Information about track
type TrackInfo struct {
	Title    string `json:"_filename"`
	Duration int
}

//YtGetTrackDetails gets information about a track
func YtGetTrackDetails(options YtOptions) (*TrackInfo, error) {
	audioData := &TrackInfo{}

	Output, err := exec.Command(
		options.Path,
		`ytsearch1:"`+options.TrackName+`"`,
		"-j",
	).Output()

	if err != nil {
		return audioData, err
	}

	err = json.Unmarshal([]byte(Output), audioData)
	if err != nil {
		return audioData, err
	}

	return audioData, nil
}

//YtSearchByName search and download youtube audio track using title
func YtSearchByName(options YtOptions) (string, error) {

	audioData, err := YtGetTrackDetails(options)
	if err != nil {
		return "", err
	}

	if audioData.Duration > 420 {
		return "", errors.New("Duration is high")
	}

	err = exec.Command(
		options.Path,
		`ytsearch1:"`+options.TrackName+`"`,
		"--extract-audio",
		"--audio-format",
		options.AudioFormat,
	).Run()

	return audioData.Title, err
}
