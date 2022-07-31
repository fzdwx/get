package msc

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fzdwx/get/pkg/ptermx"
	"github.com/fzdwx/get/pkg/utils"
	"github.com/pterm/pterm"
)

// Download music
func Download(config DownloadConfig) {
	request := getRequestFunc(config.Platform)(config.Name)

	spinnerInfo, _ := pterm.DefaultSpinner.Start("获取歌曲列表...")
	songs, songsCount, err := request.Execute()
	if err != nil {
		spinnerInfo.Fail(err.Error())
		return
	}

	if songsCount == 0 {
		spinnerInfo.Warning("没有获取到歌曲,请换个关键词再试试.")
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

func getRequestFunc(p Platform) func(name string) Request {
	switch p {
	case NetEasyP:
		return newNetEasy
	default:
		return newKuWo
	}
}

func process(s Songs) error {
	resp, err := http.Get(s.DownloadUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.OpenFile(fmt.Sprintf("%s%s", utils.NormalizeFileName(s.Name), s.EncodeType), os.O_CREATE|os.O_WRONLY, 0o777)
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
