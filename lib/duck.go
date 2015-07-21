package gfxBot

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"github.com/ajanicij/goduckgo/goduckgo"
)

func DuckSearch(query string) (string, string) {
	if query == "" {
		return "assets/img", ""
	}

	message, err := goduckgo.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return "assets/img", ""
	}

	if message != nil && message.RelatedTopics != nil && len(message.RelatedTopics) != 0  {

		filename := strings.Trim(strings.TrimSpace(query), " ")
		caption := message.RelatedTopics[0].Text
		url := strings.TrimSpace(message.RelatedTopics[0].Icon.URL)

		if url != "" {
		extensionPos := strings.LastIndex(url, ".")
		extension := url[extensionPos:len(url)]
		filename = fmt.Sprint("assets/", filename, extension)

		file, _ := os.Create(filename)
		defer file.Close()

		resp, _ := http.Get(url)
		defer resp.Body.Close()

		io.Copy(file, resp.Body)

		return filename, caption
		} else {
			return "assets/img.jpg", ""
		}
	} else {
		return "assets/img.jpg", ""
	}
}
