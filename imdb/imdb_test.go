package imdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestMovieClient_GetMovieInfo(t *testing.T) {
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
		want    MovieInfo
		wantErr bool
	}{
		{
			"MovieInfo Testing",
			fields{
				&http.Client{},
			},
			args{
				"Now you see me",
			},
			MovieInfo{
				Name:        "Now You See Me (2013)",
				Description: "An F.B.I. Agent and an Interpol Detective track a team of illusionists who pull off bank heists during their performances, and reward their audiences with the money.",
				Rating:      "7.3/10",
				PosterLink:  "https://www.imdb.com/title/tt1670345/mediaviewer/rm1351393536?ref_=tt_ov_i",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &MovieClient{
				Client: tt.fields.Client,
			}
			got, err := mc.GetMovieInfo(tt.args.inputName)
			if (err != nil) != tt.wantErr {
				t.Errorf("MovieClient.GetMovieInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovieClient.GetMovieInfo() = \n%v\n, want \n%v\n", got, tt.want)
			}
		})
	}
}
