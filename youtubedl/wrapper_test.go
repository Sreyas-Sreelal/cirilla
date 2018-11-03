package youtubedl

import (
	"reflect"
	"testing"
)

func TestYtGetTrackDetails(t *testing.T) {

	type args struct {
		options YtOptions
	}
	tests := []struct {
		name    string
		args    args
		want    *TrackInfo
		wantErr bool
	}{
		{
			"Testing by Name",
			args{
				YtOptions{
					"youtube-dl",
					"nightcore sys stand by you male",
					"mp3",
					"",
				},
			},
			&TrackInfo{
				"Nightcore -  Stand By You (Male)",
				170,
			},
			false,
		},
		{
			"Testing by Videoid",
			args{
				YtOptions{
					"youtube-dl",
					"",
					"mp3",
					"tuIkKaZPGOU",
				},
			},
			&TrackInfo{
				"Nightcore -  Stand By You (Male)",
				170,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := YtGetTrackDetails(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("YtGetTrackDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Logf("Fulltitle is %s ", got.Fulltitle)
				t.Errorf("YtGetTrackDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYtSearchByName(t *testing.T) {
	type args struct {
		options YtOptions
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Testing song Nightcore -  Stand By You (Male)",
			args{
				YtOptions{
					"youtube-dl",
					"nightcore sys stand by you male",
					"mp3",
					"",
				},
			},
			"Nightcore -  Stand By You (Male)",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := YtSearchByName(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("YtSearchByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("YtSearchByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYtExtractAudioFromID(t *testing.T) {
	type args struct {
		options YtOptions
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Testing song Nightcore -  Stand By You (Male)",
			args{
				YtOptions{
					"youtube-dl",
					"",
					"mp3",
					"tuIkKaZPGOU",
				},
			},
			"Nightcore -  Stand By You (Male)",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := YtExtractAudioFromID(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("YtExtractAudioFromID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("YtExtractAudioFromID() = %v, want %v", got, tt.want)
			}
		})
	}
}
