package msc

import (
	"fmt"
	"github.com/fzdwx/get/pkg"
	"github.com/fzdwx/get/pkg/ptermx"
	"github.com/pterm/pterm"
	"io"
	"net/http"
	"os"
)

// Download music
func Download(name string) {
	easy := newNetEasy(name)

	spinnerInfo, _ := pterm.DefaultSpinner.Start("获取歌曲列表...")
	songs, songsCount, err := easy.execute()
	if err != nil {
		spinnerInfo.Fail(err.Error())
		return
	}
	spinnerInfo.Info(fmt.Sprintf("总共获取到了%d条歌曲", songsCount))

	m, options := songsToMap(songs)

	selected, _ := pterm.DefaultInteractiveSelect.
		WithDefaultText("请选择你要下载的歌曲").
		WithMaxHeight(10).
		WithOptions(options).
		Show()

	pterm.Info.Printfln("下载: %s", selected)

	song, ok := m[selected]
	if !ok {
		return
	}

	process(song)
}

func process(s Songs) error {
	resp, err := http.Get(s.DownloadUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.OpenFile(fmt.Sprintf("%s.%s", pkg.NormalizeFileName(s.Name), s.EncodeType), os.O_CREATE|os.O_WRONLY, 777)
	if err != nil {
		return err
	}
	defer file.Close()

	p, err := pterm.DefaultProgressbar.
		WithTotal(int(resp.ContentLength)).
		WithTitle("downloading").
		Start()
	if err != nil {
		return err
	}

	_, err = io.Copy(io.MultiWriter(file, ptermx.NewProgressWriter(p)), resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func songsToMap(songs []Songs) (map[string]Songs, []string) {
	m := make(map[string]Songs, len(songs))
	var s []string

	for i, song := range songs {
		key := song.Prompt(i)
		m[key] = song
		s = append(s, key)
	}

	return m, s
}
