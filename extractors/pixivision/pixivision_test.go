package pixivision

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
				URL:   "https://www.pixivision.net/zh/a/3271",
				Title: "Don't ask me to choose! Tiny Breasts VS Huge Breasts",
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
