package gfxBot

import (
	"errors"
	"fmt"
	"os"

	"github.com/gloob/go-instagram/instagram"
)

// Instagram implements the gfxbot instagram service.
type Instagram struct {
	APIKey    string
	APISecret string
	APIToken  string
	Client    *instagram.Client
	Service   string
}

// NewInstagram returns an initialized Instagram service.
func NewInstagram(apiKey string, apiSecret string, apiToken string) *Instagram {
	return &Instagram{
		APIKey:    apiKey,
		APISecret: apiSecret,
		APIToken:  apiToken,
		Client:    instagram.NewClient(nil),
		Service:   "Instagram",
	}
}

// Search into Instagram service for a given string.
func (i *Instagram) Search(q string) (*Image, error) {

	i.Client.ClientID = i.APIKey
	i.Client.ClientSecret = i.APISecret
	i.Client.AccessToken = i.APIToken

	if q == "" {
		return nil, errors.New("Wrong syntax!. Syntax is: /instagram <query_string>. Example: /instagram new trends")
	}

	var tagName string
	tags, _, err := i.Client.Tags.Search(q)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return nil, err
	}

	// Select the first one if exists or the query string.
	if len(tags) > 0 {
		tagName = tags[0].Name
	} else {
		tagName = q
	}

	/*
		// Iterate over the slice of tags.
		for _, t := range tags {
			tagName = t.Name
		}
	*/

	// Obtain information about the selected tag.
	tagInfo, err := i.Client.Tags.Get(tagName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return nil, err
	}

	opt := &instagram.Parameters{
		MinID: "0",
		MaxID: "999999999",
		Count: 10,
	}
	medias, _, err := i.Client.Tags.RecentMedia(tagInfo.Name, opt)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return nil, err
	}

	var url, caption string
	for _, media := range medias {
		url = media.Images.StandardResolution.URL
		caption = media.Caption.Text
	}
	fmt.Fprintf(os.Stdout, "%v\n", url)

	if url != "" {
		img, _ := NewImage(url, caption)
		return img, nil
	}

	return nil, errors.New("Instagram: No images found.")
}
