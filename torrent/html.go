package torrent

import (
	"github.com/PuerkitoBio/goquery"
)

//GetHTMLDoc returns html document
func (pb *PbClient) GetHTMLDoc(url string) (*goquery.Document, error) {
	res, err := pb.Client.Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	return doc, err
}
