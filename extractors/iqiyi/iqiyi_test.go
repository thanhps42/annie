package iqiyi

import (
	"testing"

	"github.com/thanhps42/annie/config"
	"github.com/thanhps42/annie/test"
)

func TestDownload(t *testing.T) {
	config.InfoOnly = true
	config.RetryTimes = 100
	tests := []struct {
		name string
		args test.Args
	}{
		{
			name: "normal test",
			args: test.Args{
				URL:     "http://www.iqiyi.com/v_19rrbdmaj0.html",
				Title:   "新一轮降水将至 冷空气影响中东部地区-资讯-高清正版视频在线观看–爱奇艺",
				Size:    2952228,
				Quality: "896x504",
			},
		},
		{
			name: "title test 1",
			args: test.Args{
				URL:     "http://www.iqiyi.com/v_19rqy2z83w.html",
				Title:   "收了创意视频2018：58天环球飞行记",
				Size:    76186786,
				Quality: "1920x1080",
			},
		},
		{
			name: "curid test 1",
			args: test.Args{
				URL:     "https://www.iqiyi.com/v_19rro0jdls.html#curid=350289100_6e6601aae889d0b1004586a52027c321",
				Title:   "Shawn Mendes - Never Be Alone",
				Size:    79921894,
				Quality: "1920x800",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Extract(tt.args.URL)
			test.CheckError(t, err)
			test.Check(t, tt.args, data[0])
		})
	}
}
