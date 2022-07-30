package msc

import (
	"testing"
)

/*

{
"name":"稻香 (钢琴版) [原唱: 周杰伦]",
"artists":[{}],
"downloadUrl":"http://m701.music.126.net/20220730162101/f870cf2d08e9b23449dcaa6408048561/jdymusic/obj/wo3DlMOGwrbDjj7DisKw/14364323356/5c3e/a89c/b095/2c154c2ce5ed4487cefbfa1abc2c0f57.mp3",
"size":3138477,
"encodeType":"mp3"
}
*/
func Test_process(t *testing.T) {
	songs := Songs{
		Name:        "稻香 (钢琴版) [原唱: 周杰伦]",
		Artists:     []Artists{},
		DownloadUrl: "http://m701.music.126.net/20220730162101/f870cf2d08e9b23449dcaa6408048561/jdymusic/obj/wo3DlMOGwrbDjj7DisKw/14364323356/5c3e/a89c/b095/2c154c2ce5ed4487cefbfa1abc2c0f57.mp3",
		Size:        3138477,
		EncodeType:  "mp3",
	}
	process(songs)
}
