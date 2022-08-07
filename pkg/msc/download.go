package msc

import (
	"fmt"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/fzdwx/get/pkg/ptermx"
	"github.com/fzdwx/get/pkg/utils"
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/progress"
	"github.com/fzdwx/infinite/components/selection/multiselect"
	"github.com/fzdwx/infinite/components/spinner"
	"github.com/fzdwx/infinite/style"
	"io"
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
		spinner.WithPrompt("  获取歌曲列表..."),
		spinner.WithFunc(func(spinner *spinner.Spinner) {
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
		}),
	).Display()

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
	err := progress.NewGroupWithCount(len(selectedSongs)).AppendRunner(func(progress *components.Progress) func() {
		return func() {
			s := selectedSongs[progress.Id-1]
			resp, err := http.Get(s.DownloadUrl)
			defer resp.Body.Close()
			if err != nil {
				// todo handle err
				return
			}

			file, err := os.OpenFile(fmt.Sprintf("%s%s", utils.NormalizeFileName(s.Name), s.EncodeType), os.O_CREATE|os.O_WRONLY, 0o777)
			defer file.Close()
			if err != nil {
				// todo handle err
				return
			}

			progress.
				WithTotal(resp.ContentLength).
				WithPercentAgeFunc(func(total int64, current int64, percent float64) string {
					return fmt.Sprintf(" %d/%d", current, total)
				})

			_, err = io.Copy(io.MultiWriter(file, ptermx.NewProgressWriter(progress)), resp.Body)
			if err != nil {
				// todo handle err
			}
		}
	}).Display()

	if err != nil {
		fmt.Println(err)
	}

	return nil
}
