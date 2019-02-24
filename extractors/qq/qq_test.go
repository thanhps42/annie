package qq

import (
	"testing"

	"github.com/thanhps42/annie/config"
	"github.com/thanhps42/annie/test"
)

func TestDownload(t *testing.T) {
	config.InfoOnly = true
	config.RetryTimes = 10
	tests := []struct {
		name string
		args test.Args
	}{
		{
			name: "normal test",
			args: test.Args{
				URL:     "https://v.qq.com/x/page/n0687peq62x.html",
				Title:   "世界杯第一期：100秒速成！“伪球迷”世界杯生存指南",
				Size:    23759683,
				Quality: "蓝光;(1080P)",
			},
		},
		// {
		// 	name: "movie and vid test",
		// 	args: test.Args{
		// 		URL:     "https://v.qq.com/x/cover/e5qmd3z5jr0uigk.html",
		// 		Title:   "赌侠（粤语版）",
		// 		Size:    1046910811,
		// 		Quality: "超清;(720P)",
		// 	},
		// },
		{
			name: "single part test",
			args: test.Args{
				URL:     "https://v.qq.com/iframe/player.html?vid=v0739eolv38",
				Title:   "PGI国际邀请赛，FPP第四局，OMG强势吃鸡，全场观众高喊OMG",
				Size:    10714773,
				Quality: "高清;(480P)",
			},
		},
		{
			name: "fmt ID test",
			args: test.Args{
				URL:     "https://v.qq.com/x/cover/2aya3ibdmft6vdw/e0765r4mwcr.html",
				Title:   "《卡路里》出圈！妖娆男子教学广场舞版，大妈表情亮了！",
				Size:    14112979,
				Quality: "超清;(720P)",
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
