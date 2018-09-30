package imdb

import (
	"reflect"
	"testing"
)

func TestMovieClient_GetMovieInfo(t *testing.T) {

	type args struct {
		inputName string
	}
	tests := []struct {
		name string

		args    args
		want    MovieInfo
		wantErr bool
	}{
		{
			"Testing Starwars",
			args{"Solo: A Star Wars Story"},
			MovieInfo{
				"Solo: A Star Wars Story (2018)",
				"7.1/10",
				"A Star Wars Story (2018) 7.1/10 During an adventure into the criminal underworld, Han Solo meets his future co-pilot Chewbacca and encounters Lando Calrissian years before joining the Rebellion.",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := GetNewClient()
			got, err := mc.GetMovieInfo(tt.args.inputName)
			if (err != nil) != tt.wantErr {
				t.Errorf("MovieClient.GetMovieInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovieClient.GetMovieInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
