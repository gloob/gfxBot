package gfxBot

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/masci/flickr"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

type Flickr struct {
	ApiKey    string
	ApiSecret string
	Client    *flickr.FlickrClient
}

type Photo struct {
	Name     string `xml:",chardata"`
	Id       int    `xml:"id,attr"`
	Owner    string `xml:"owner,attr"`
	Secret   string `xml:"secret,attr"`
	Server   int    `xml:"server,attr"`
	Farm     int    `xml:"farm,attr"`
	Title    string `xml:"title,attr"`
	IsPublic bool   `xml:"ispublic,attr"`
	IsFriend bool   `xml:"isfriend,attr"`
	IsFamily bool   `xml:"isfamily,attr"`
}

type Photos struct {
	XMLName xml.Name `xml:"photos"`
	Photos  []Photo  `xml:"photo"`
}

func NewFlickr(apiKey string, apiSecret string) *Flickr {
	return &Flickr{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
		Client:    flickr.NewFlickrClient(apiKey, apiSecret),
	}
}

func (f *Flickr) Search(q string) (*Image, error) {
	if q == "" {
		return nil, errors.New("Trying to make a query without parameters")
	}
	f.Client.Init()
	f.Client.Args.Set("api_key", f.ApiKey)
	f.Client.Args.Set("method", "flickr.photos.search")
	f.Client.Args.Set("text", strings.Trim(q, " "))
	f.Client.Args.Set("format", "rest")

	response := &flickr.BasicResponse{}
	fmt.Println(f.Client.GetUrl())
	err := flickr.DoGet(f.Client, response)
	if err != nil {
		return nil, err
	}

	var p Photos
	xml.Unmarshal([]byte(response.Extra), &p)

	if len(p.Photos) <= 0 {
		return nil, errors.New("No results")
	}

	idx := rand.Intn(len(p.Photos))
	photo := p.Photos[idx]

	img, _ := NewImage(".jpg", photo.Title)
	defer img.Close()

	resp, err := http.Get(photo.buildUrl())
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.Body)
	defer resp.Body.Close()

	io.Copy(img, resp.Body)

	return img, nil
}

func (p Photo) buildUrl() string {
	return fmt.Sprintf("https://farm%d.staticflickr.com/%d/%d_%s_z.jpg",
		p.Farm,
		p.Server,
		p.Id,
		p.Secret)
}
