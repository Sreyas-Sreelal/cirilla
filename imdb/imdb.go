package imdb

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strings"
)

//MovieClient represent http client that access imdb site
type MovieClient struct {
	Client *http.Client
}

//MovieInfo stores information about movie fetched from imdb
type MovieInfo struct {
	Name        string
	Rating      string
	Description string
	PosterLink  string
}

//GetNewClient creates new instance of MovieClient
func GetNewClient() *MovieClient {
	var mc *MovieClient
	mc = new(MovieClient)
	mc.Client = &http.Client{}
	return mc
}

//SearchTitleURL returns result url after searching
func (mc *MovieClient) SearchTitleURL(inputName string) (string, error) {
	document, err := mc.GetHTMLDoc(fmt.Sprintf("https://www.imdb.com/find?q=%s&s=tt", inputName))
	if err != nil {
		log.Printf("[Error] Failed in geting HtmlDoc for input %s\n %q", inputName, err)
		return "", err
	}

	url, success := document.Find(".result_text").First().Find("a").Attr("href")

	if !success {
		return "", errors.New(fmt.Sprintf("Failed getting href attribute value for input %s", inputName))
	}
	log.Printf("Url for %s is %s ", inputName, url)

	return url, nil
}

//GetMovieName fetchs movie name
func (mc *MovieClient) GetMovieName(document *goquery.Document) (string, error) {
	movieNameSelector := document.Find(".title_wrapper").First().Find("h1")
	movieName := movieNameSelector.Text()
	if movieName == "" {
		return "", errors.New(fmt.Sprintf("Failed getting movie name for %q", document))
	}

	movieName = strings.Replace(strings.TrimSpace(movieName), "\u00a0", " ", -1)

	return movieName, nil
}

//GetMovieRating fetchs movie rating
func (mc *MovieClient) GetMovieRating(document *goquery.Document) (string, error) {
	movieRatingSelector := document.Find(".ratingValue").First()
	movieRating := movieRatingSelector.Text()
	if movieRating == "" {
		return "", errors.New(fmt.Sprintf("Failed getting movie rating for %q", document))
	}

	movieRating = strings.TrimSpace(movieRating)

	return movieRating, nil
}

//GetMovieDescription fetchs movie summary
func (mc *MovieClient) GetMovieDescription(document *goquery.Document) (string, error) {
	movieDescriptionSelector := document.Find(".summary_text").First()
	movieDescription := movieDescriptionSelector.Text()
	if movieDescription == "" {
		return "", errors.New(fmt.Sprintf("Failed getting movie description for %q", document))
	}

	movieDescription = strings.TrimSpace(movieDescription)

	return movieDescription, nil
}

//GetMoviePoster fetchs movie name
func (mc *MovieClient) GetMoviePoster(document *goquery.Document) (string, error) {
	url, success := document.Find(".poster").First().Find("a").Attr("href")

	if !success {
		return "", errors.New(fmt.Sprintf("Failed getting movie poster for %q", document))
	}

	url = "https://www.imdb.com" + url

	return url, nil
}

//GetMovieInfo fetches information about a movie
func (mc *MovieClient) GetMovieInfo(inputName string) (MovieInfo, error) {
	inputName = strings.Replace(inputName, " ", "+", len(inputName))
	url, err := mc.SearchTitleURL(inputName)
	if err != nil {
		return MovieInfo{}, err
	}

	document, err := mc.GetHTMLDoc(fmt.Sprintf("https://www.imdb.com%s", url))
	if err != nil {
		return MovieInfo{}, err
	}

	movieName, err := mc.GetMovieName(document)
	if err != nil {
		return MovieInfo{}, err
	}

	movieRating, err := mc.GetMovieRating(document)
	if err != nil {
		return MovieInfo{}, err
	}

	movieDescription, err := mc.GetMovieDescription(document)
	if err != nil {
		return MovieInfo{}, err
	}

	moviePosterLink, err := mc.GetMoviePoster(document)
	if err != nil {
		return MovieInfo{}, err
	}

	return MovieInfo{movieName, movieRating, movieDescription, moviePosterLink}, nil

}
