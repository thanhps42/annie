package yinyuetai

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/thanhps42/annie/downloader"
	"github.com/thanhps42/annie/request"
	"github.com/thanhps42/annie/utils"
)

const yinyuetaiAPI = "https://ext.yinyuetai.com/main/"

const (
	actionGetMvInfo = "get-h-mv-info"
)

func genAPI(action string, param string) string {
	return fmt.Sprintf("%s%s?json=true&%s", yinyuetaiAPI, action, param)
}

// Extract is the main function for extracting data
func Extract(url string) ([]downloader.Data, error) {
	vid := utils.MatchOneOf(
		url,
		`https?://v.yinyuetai.com/video/(\d+)(?:\?vid=\d+)?`,
		`https?://v.yinyuetai.com/video/h5/(\d+)(?:\?vid=\d+)?`,
		`https?://m2.yinyuetai.com/video.html\?id=(\d+)`,
	)
	if vid == nil {
		return downloader.EmptyList, errors.New("invalid url for yinyuetai")
	}
	params := fmt.Sprintf("videoId=%s", vid[1])
	// generate api url
	apiUrl := genAPI(actionGetMvInfo, params)
	var err error
	html, err := request.Get(apiUrl, url, nil)
	if err != nil {
		return downloader.EmptyList, err
	}
	// parse yinyuetai data
	data := yinyuetaiMvData{}
	json.Unmarshal([]byte(html), &data)
	// handle api error
	if data.Error {
		return downloader.EmptyList, errors.New(data.Message)
	}
	if data.VideoInfo.CoreVideoInfo.Error {
		return downloader.EmptyList, errors.New(data.VideoInfo.CoreVideoInfo.ErrorMsg)
	}
	title := data.VideoInfo.CoreVideoInfo.VideoName
	streams := map[string]downloader.Stream{}
	// set streams
	for _, model := range data.VideoInfo.CoreVideoInfo.VideoURLModels {
		urlData := downloader.URL{
			URL:  model.VideoURL,
			Size: model.FileSize,
			Ext:  "mp4",
		}
		streams[model.QualityLevel] = downloader.Stream{
			URLs:    []downloader.URL{urlData},
			Size:    model.FileSize,
			Quality: model.QualityLevelName,
		}
	}
	return []downloader.Data{
		{
			Site:    "音悦台 yinyuetai.com",
			Title:   title,
			Type:    "video",
			Streams: streams,
			URL:     url,
		},
	}, nil
}
