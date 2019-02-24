package bcy

import (
	"encoding/json"
	"strings"

	"github.com/thanhps42/annie/downloader"
	"github.com/thanhps42/annie/parser"
	"github.com/thanhps42/annie/request"
	"github.com/thanhps42/annie/utils"
)

type bcyData struct {
	Detail struct {
		PostData struct {
			Multi []struct {
				OriginalPath string `json:"original_path"`
			} `json:"multi"`
		} `json:"post_data"`
	} `json:"detail"`
}

// Extract is the main function for extracting data
func Extract(url string) ([]downloader.Data, error) {
	var err error
	html, err := request.Get(url, url, nil)
	if err != nil {
		return downloader.EmptyList, err
	}

	// parse json data
	rep := strings.NewReplacer(`\"`, `"`, `\\`, `\`)
	jsonString := rep.Replace(utils.MatchOneOf(html, `JSON.parse\("(.+?)"\);`)[1])
	var data bcyData
	json.Unmarshal([]byte(jsonString), &data)

	doc, err := parser.GetDoc(html)
	if err != nil {
		return downloader.EmptyList, err
	}
	title := strings.Replace(parser.Title(doc), " - 半次元 banciyuan - ACG爱好者社区", "", -1)

	urls := make([]downloader.URL, 0, len(data.Detail.PostData.Multi))
	var totalSize int64
	for _, img := range data.Detail.PostData.Multi {
		size, err := request.Size(img.OriginalPath, url)
		if err != nil {
			return downloader.EmptyList, err
		}
		totalSize += size
		_, ext, err := utils.GetNameAndExt(img.OriginalPath)
		if err != nil {
			return downloader.EmptyList, err
		}
		urls = append(urls, downloader.URL{
			URL:  img.OriginalPath,
			Size: size,
			Ext:  ext,
		})
	}
	streams := map[string]downloader.Stream{
		"default": {
			URLs: urls,
			Size: totalSize,
		},
	}
	return []downloader.Data{
		{
			Site:    "半次元 bcy.net",
			Title:   title,
			Type:    "image",
			Streams: streams,
			URL:     url,
		},
	}, nil
}
