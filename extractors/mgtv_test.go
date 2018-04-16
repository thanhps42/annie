package extractors

import (
	"testing"

	"github.com/iawia002/annie/config"
	"github.com/iawia002/annie/test"
)

func TestMgtv(t *testing.T) {
	config.InfoOnly = true
	tests := []struct {
		name string
		args test.Args
	}{
		{
			name: "normal test",
			args: test.Args{
				URL:     "https://www.mgtv.com/b/322712/4317248.html",
				Title:   "我是大侦探 先导片：何炅吴磊邓伦穿越破案",
				Size:    86169236,
				Quality: "超清",
			},
		},
		{
			name: "normal test",
			args: test.Args{
				URL:     "https://www.mgtv.com/b/308703/4197072.html",
				Title:   "芒果捞星闻 2017 诺一为爷爷和姥爷做翻译超萌",
				Size:    6486376,
				Quality: "超清",
			},
		},
		{
			name: "vip test",
			args: test.Args{
				URL:     "https://www.mgtv.com/b/322865/4352046.html",
				Title:   "向往的生活 第二季 先导片：何炅黄磊回归质朴生活",
				Size:    425531232,
				Quality: "超清",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := Mgtv(tt.args.URL)
			test.Check(t, tt.args, data)
		})
	}
}
