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
	VideoID     string
}

//TrackInfo holds Information about track
type TrackInfo struct {
	Fulltitle string `json:"fulltitle"`
	Duration  int
}

//YtGetTrackDetails gets information about a track
func YtGetTrackDetails(options YtOptions) (*TrackInfo, error) {
	audioData := &TrackInfo{}
	var specifier string

	if options.TrackName != "" {
		specifier = `ytsearch1:"` + options.TrackName + `"`
	} else if options.VideoID != "" {
		specifier = options.VideoID
	} else {
		return audioData, errors.New("No TrackName or VideoID passed")
	}

	Output, err := exec.Command(
		options.Path,
		specifier,
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

//YtSearchByName search and download youtube audio track using Fulltitle
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
		"--output",
		audioData.Fulltitle+"."+options.AudioFormat,
	).Run()

	return audioData.Fulltitle, err
}

//YtExtractAudioFromID extracts audio and onverts it into desired format
func YtExtractAudioFromID(options YtOptions) (string, error) {
	audioData, err := YtGetTrackDetails(options)
	if err != nil {
		return "", err
	}
	if audioData.Duration > 420 {
		return "", errors.New("Duration is high")
	}

	err = exec.Command(
		options.Path,
		options.VideoID,
		"--extract-audio",
		"--audio-format",
		options.AudioFormat,
		"--output",
		audioData.Fulltitle+"."+options.AudioFormat,
	).Run()

	return audioData.Fulltitle, err
}
