package msc

import (
	"fmt"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/fzdwx/get/pkg/utils"
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/progress"
	"github.com/fzdwx/infinite/components/selection/multiselect"
	"github.com/fzdwx/infinite/components/spinner"
	"github.com/fzdwx/infinite/style"
	"net/http"
	"os"
)

// Download music
func Download(config DownloadConfig) {
	request := getRequestFunc(config.Platform)(config.Name)

	var (
		songs      []Songs
		songsCount int
		err        error
	)

	_ = inf.NewSpinner(
		spinner.WithPrompt("  获取歌曲列表...")).
		Display(func(spinner *spinner.Spinner) {
			songs, songsCount, err = request.Execute()
			if err != nil {
				spinner.Finish(err.Error())
				return
			}

			if songsCount == 0 {
				spinner.Finish("没有获取到歌曲,请换个关键词再试试.")
				return
			}

			spinner.Finish(fmt.Sprintf("总共获取到了%d条歌曲", songsCount))
		})

	options := slice.Map[Songs, string](songs, func(index int, item Songs) string {
		return item.Prompt(index)
	})

	input := components.NewInput()
	input.Prompt = "Filtering: "
	input.PromptStyle = style.New().Bold().Italic().Fg(color.LightBlue)
	ints, err := inf.NewMultiSelect(options,
		multiselect.WithPageSize(10),
		multiselect.WithFilterInput(input),
	).Display("请选择你要下载的歌曲:")

	selectedSongs := slice.Map[int, Songs](ints, func(index int, item int) Songs {
		return songs[item]
	})

	process(selectedSongs)
}

func getRequestFunc(p Platform) func(name string) Request {
	switch p {
	case NetEasyP:
		return newNetEasy
	default:
		return newKuWo
	}
}

func process(selectedSongs []Songs) error {
	err := progress.NewGroupWithCount(len(selectedSongs)).AppendRunner(func(pro *components.Progress) func() {
		s := selectedSongs[pro.Id-1]
		return func() {

			resp, err := http.Get(s.DownloadUrl)
			if err != nil {
				resp.Body.Close()
				pro.WithDoneView(func() string {
					return fmt.Sprintf("get error: %s", err)
				})
				return
			}

			pro.WithTotal(resp.ContentLength)

			fileName := fmt.Sprintf("%s%s", utils.NormalizeFileName(s.Name), s.EncodeType)
			dest, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0o777)
			if err != nil {
				dest.Close()
				pro.WithDoneView(func() string {
					return fmt.Sprintf("open dest error: %s", err)
				})
				return
			}

			defer resp.Body.Close()
			defer dest.Close()

			pro.WithDoneView(func() string {
				return fmt.Sprintf("%s download success", fileName)
			})

			_, err = progress.StartTransfer(resp.Body, dest, pro)
			if err != nil {
				pro.WithDoneView(func() string {
					return fmt.Sprintf("transfer error: %s", err)
				})
			}
		}
	}).Display()

	if err != nil {
		fmt.Println(err)
	}

	return nil
}
