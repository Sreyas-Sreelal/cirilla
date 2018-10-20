package torrent

import (
	"net/http"
	"testing"
)

func TestPbClient_GetTorrentInfo(t *testing.T) {
	type fields struct {
		Client *http.Client
	}
	type args struct {
		inputName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    PbPageInfo
		wantErr bool
	}{
		{
			"TorrentInfo Testing",
			fields{
				&http.Client{},
			},
			args{
				"Witcher 3",
			},
			PbPageInfo{
				Name:      "Witcher3",
				MagnetURL: "magnet:?xt=urn:btih:f6921cf841c1d8a6b1233eac6034303e6f40f4b5&dn=The+Witcher+3+Wild+Hunt+Game+of+the+Year+Edition+PROPER-GOG&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969&tr=udp%3A%2F%2Fzer0day.ch%3A1337&tr=udp%3A%2F%2Fopen.demonii.com%3A1337&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969&tr=udp%3A%2F%2Fexodus.desync.com%3A6969",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pb := &PbClient{
				Client: tt.fields.Client,
			}
			got, err := pb.GetTorrentInfo(tt.args.inputName)
			if (err != nil) != tt.wantErr {
				t.Errorf("PbClient.GetTorrentInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.MagnetURL != tt.want.MagnetURL {
				t.Errorf("PbClient.GetTorrentInfo() = %s, want %s", got.MagnetURL, tt.want.MagnetURL)
			}
		})
	}
}
