package youtube

import (
	"testing"

	"github.com/thanhps42/annie/config"
	"github.com/thanhps42/annie/downloader"
	"github.com/thanhps42/annie/test"
)

func TestYoutube(t *testing.T) {
	config.InfoOnly = true
	config.ThreadNumber = 9
	tests := []struct {
		name     string
		args     test.Args
		playlist bool
		stream2  bool
	}{
		{
			name: "normal test",
			args: test.Args{
				URL:     "https://www.youtube.com/watch?v=Gnbch2osEeo",
				Title:   "Multifandom Mashup 2017",
				Quality: `720p video/mp4; codecs="avc1.4d401f"`,
			},
		},
		{
			name: "normal test",
			args: test.Args{
				URL:     "https://youtu.be/z8eFzkfto2w",
				Title:   "Circle Of Love | Rudy Mancuso",
				Size:    33060990,
				Quality: `1080p video/webm; codecs="vp9"`,
			},
		},
		{
			name: "normal test",
			args: test.Args{
				URL:     "https://www.youtube.com/watch?v=ASPku-eAZYs",
				Title:   "怪獸與葛林戴華德的罪行 | HD首版電影預告大首播 (Fantastic Beasts: The Crimes of Grindelwald)",
				Quality: `1080p60 video/mp4; codecs="avc1.64002a"`,
			},
		},
		{
			name: "playlist test",
			args: test.Args{
				URL:     "https://www.youtube.com/watch?v=x2nKigmfzLQ&list=PLfyyPldyYFapPfxZOay2AoCCaNE1Ezd4_",
				Title:   "【仔細看看】《羅根》黑白版-漫畫電影新高峰?! - 超粒方",
				Size:    192310804,
				Quality: `1080p video/webm; codecs="vp9"`,
			},
			playlist: true,
		},
		{
			name: "url_encoded_fmt_stream_map test",
			args: test.Args{
				URL:     "https://youtu.be/DNaOZovrSVo",
				Title:   "QNAP Case Study - Scorptec",
				Size:    25418418,
				Quality: `1080p video/mp4; codecs="avc1.640028"`,
			},
		},
		{
			name: "stream 404 test 1",
			args: test.Args{
				URL:     "https://www.youtube.com/watch?v=MRJ8NnUXacY",
				Title:   "FreeFileSync: Mirror Synchronization",
				Quality: `1080p60 video/mp4; codecs="avc1.64002a"`,
			},
		},
		{
			name: "stream 404 test 2",
			args: test.Args{
				URL:     "https://www.youtube.com/watch?v=MRJ8NnUXacY",
				Title:   "FreeFileSync: Mirror Synchronization",
				Quality: `hd720 video/mp4; codecs="avc1.64001F, mp4a.40.2"`,
			},
			stream2: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				data []downloader.Data
				err  error
			)
			if tt.stream2 {
				config.YouTubeStream2 = true
			} else {
				config.YouTubeStream2 = false
			}
			if tt.playlist {
				// playlist mode
				config.Playlist = true
				_, err = Extract(tt.args.URL)
				test.CheckError(t, err)
			} else {
				config.Playlist = false
				data, err = Extract(tt.args.URL)
				test.CheckError(t, err)
				test.Check(t, tt.args, data[0])
			}
		})
	}
}
