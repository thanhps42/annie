package facebook

import (
	"fmt"
	"strings"

	"github.com/thanhps42/annie/downloader"
	"github.com/thanhps42/annie/request"
	"github.com/thanhps42/annie/utils"
)

// Extract is the main function for extracting data
func Extract(url string) ([]downloader.Data, error) {
	var err error
	html, err := request.Get(url, url, nil)
	if err != nil {
		return downloader.EmptyList, err
	}
	title := utils.MatchOneOf(html, `<title id="pageTitle">((.+)(\n(.+))?)</title>`)[1]
	title = strings.Replace(title, "\n", " ", -1)

	streams := map[string]downloader.Stream{}
	for _, quality := range []string{"sd", "hd"} {
		u := utils.MatchOneOf(
			html, fmt.Sprintf(`%s_src:"(.+?)"`, quality),
		)[1]
		size, err := request.Size(u, url)
		if err != nil {
			return downloader.EmptyList, err
		}
		urlData := downloader.URL{
			URL:  u,
			Size: size,
			Ext:  "mp4",
		}
		streams[quality] = downloader.Stream{
			URLs:    []downloader.URL{urlData},
			Size:    size,
			Quality: quality,
		}
	}

	return []downloader.Data{
		{
			Site:    "Facebook facebook.com",
			Title:   title,
			Type:    "video",
			Streams: streams,
			URL:     url,
		},
	}, nil
}
