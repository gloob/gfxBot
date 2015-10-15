package gfxBot

import (
	"errors"
	"github.com/ajanicij/goduckgo/goduckgo"
	"io"
	"net/http"
	"strings"
)

func DuckSearch(q string) (*Image, error) {
	if q == "" {
		return nil, errors.New("DuckSearch: called with an empty query string.")
	}

	message, err := goduckgo.Query(q)
	if err != nil {
		return nil, err
	}

	if message != nil && message.RelatedTopics != nil && len(message.RelatedTopics) != 0 {

		caption := message.RelatedTopics[0].Text
		url := strings.TrimSpace(message.RelatedTopics[0].Icon.URL)

		if url != "" {
			extensionPos := strings.LastIndex(url, ".")
			extension := url[extensionPos:len(url)]

			img, _ := NewImage(extension, caption)
			defer img.Close()

			resp, _ := http.Get(url)
			defer resp.Body.Close()

			io.Copy(img, resp.Body)

			return img, nil
		} else {
			return nil, errors.New("DuckSearch: response from Duck service contains an empty URL.")
		}
	} else {
		return nil, errors.New("DuckSearch: called with an empty query string.")
	}
}
