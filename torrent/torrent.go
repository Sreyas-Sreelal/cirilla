package torrent

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strings"
)

//PbClient represent http access to PirateBay
type PbClient struct {
	Client *http.Client
}

//PbPageInfo holds torrent information
type PbPageInfo struct {
	Name        string
	MagnetURL   string
	Description string
	URL         string
}

//GetNewClient creates new instance of PbClient
func GetNewClient() *PbClient {
	var pb *PbClient
	pb = new(PbClient)
	pb.Client = &http.Client{}
	return pb
}

//GetProxy grabs piratebay proxies
func (pb *PbClient) GetProxy() (string, error) {
	var err error
	document, err := pb.GetHTMLDoc("https://thepiratebay-proxylist.se")
	if err != nil {
		return "", err
	}

	var ProxyLink string
	var success bool
	found := false

	document.Find(".url").EachWithBreak(func(i int, links *goquery.Selection) bool {
		ProxyLink, success = links.Attr("data-href")
		log.Printf("Trying %s", ProxyLink)
		_, err = pb.Client.Get(ProxyLink)
		if err == nil && success {
			log.Printf("Success %s", ProxyLink)
			found = true
			return false
		}
		return true
	})

	if !found {
		return "", errors.New("Failed to fetch proxy")
	}

	return ProxyLink, nil
}

//GetTorrentURL grabs torrent url for specific torrent
func (pb *PbClient) GetTorrentURL(inputName string) (string, error) {
	ProxyLink, err := pb.GetProxy()
	if err != nil {
		return "", err
	}

	document, err := pb.GetHTMLDoc(ProxyLink + "/s/?q=" + inputName + "&page=0&orderby=99")
	if err != nil {
		return "", err
	}

	URL, success := document.Find(".detName").First().Find("a").First().Attr("href")
	if !success {
		return "", errors.New("Failed fetching Torrent URL")
	}

	URL = ProxyLink + URL

	return URL, nil
}

//GetMagnetURL grabs magnet url for specific torrent
func (pb *PbClient) GetMagnetURL(document *goquery.Document) (string, error) {

	MagnetURL, success := document.Find(".download").First().Find("a").First().Attr("href")
	if !success {
		return "", errors.New("Failed fetching Magnet URL")
	}

	return MagnetURL, nil
}

//GetDescription grabs magnet url for specific torrent
func (pb *PbClient) GetDescription(document *goquery.Document) (string, error) {
	Description := document.Find(".nfo").First().Text()

	return Description, nil
}

//GetName grabs magnet url for specific torrent
func (pb *PbClient) GetName(document *goquery.Document) (string, error) {
	TorrentName := document.Find("Title").First().Text()

	return TorrentName, nil
}

//GetTorrentInfo gets torrent informatio from pirate bay
func (pb *PbClient) GetTorrentInfo(inputName string) (PbPageInfo, error) {
	inputName = strings.Replace(inputName, " ", "+", -1)

	URL, err := pb.GetTorrentURL(inputName)
	if err != nil {
		return PbPageInfo{}, err
	}

	document, err := pb.GetHTMLDoc(URL)
	if err != nil {
		return PbPageInfo{}, err
	}

	Name, err := pb.GetName(document)
	if err != nil {
		return PbPageInfo{}, err
	}

	MagnetURL, err := pb.GetMagnetURL(document)
	if err != nil {
		return PbPageInfo{}, err
	}

	Description, err := pb.GetDescription(document)
	if err != nil {
		return PbPageInfo{}, err
	}

	return PbPageInfo{Name, MagnetURL, Description, URL}, nil
}
