package imdb

import (
	"github.com/PuerkitoBio/goquery"
)

//GetHTMLDoc returns html document
func (mc *MovieClient) GetHTMLDoc(url string) (*goquery.Document, error) {
	res, err := mc.Client.Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	return doc, err
}
