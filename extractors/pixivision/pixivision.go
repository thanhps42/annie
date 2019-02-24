package pixivision

import (
	"github.com/thanhps42/annie/downloader"
	"github.com/thanhps42/annie/parser"
	"github.com/thanhps42/annie/request"
)

// Extract is the main function for extracting data
func Extract(url string) ([]downloader.Data, error) {
	html, err := request.Get(url, url, nil)
	if err != nil {
		return downloader.EmptyList, err
	}
	title, urls, err := parser.GetImages(url, html, "am__work__illust  ", nil)
	if err != nil {
		return downloader.EmptyList, err
	}
	streams := map[string]downloader.Stream{
		"default": {
			URLs: urls,
			Size: 0,
		},
	}

	return []downloader.Data{
		{
			Site:    "pixivision pixivision.net",
			Title:   title,
			Type:    "image",
			Streams: streams,
			URL:     url,
		},
	}, nil
}
