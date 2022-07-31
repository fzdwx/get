package msc

import (
	"fmt"
	"github.com/fzdwx/get/pkg/utils"
	"github.com/pterm/pterm"
	"sync"
)

type (
	Request interface {
		// Execute request get songs
		Execute() ([]Songs, int, error)
		// generate request url
		url() string
		// next page
		nextPage()
		// prev page
		prevPage()
	}

	SongsMapper interface {
		mapper() (*Songs, error)
		name() string
	}

	Platform int

	DownloadConfig struct {
		Name     string
		Platform Platform
	}

	Songs struct {
		// the song Name
		Name string `json:"name"`
		// the song author
		Artists []Artists `json:"artists"`
		// the song download url
		DownloadUrl string `json:"downloadUrl"`
		// the song Size(unit is byte)
		Size int64 `json:"size"`
		// the song EncodeType
		EncodeType string `json:"encodeType"`
	}

	Artists struct {
		name string
	}
)

const (
	NetEasyP Platform = 1
	KuWoP    Platform = 2
)

func (s Songs) Prompt(i int) string {
	return fmt.Sprintf("%d. %s(%s) - %s", i, s.Name, utils.FormatBytes(s.Size), utils.MappingArtName(s.Artists[0].name))
}

func collect(mappers []SongsMapper) []Songs {
	songCh := make(chan Songs)
	wg := sync.WaitGroup{}
	wg.Add(len(mappers))

	var songs []Songs
	go func() {
		for i := range mappers {
			go func(mapper SongsMapper) {
				mscSongs, err := mapper.mapper()
				if err != nil {
					wg.Done()
					pterm.Error.Printfln("download %s fail", mapper.name())
				}

				songCh <- *mscSongs
			}(mappers[i])
		}
	}()

	go func() {
		for {
			select {
			case s, ok := <-songCh:
				if ok {
					wg.Done()
					songs = append(songs, s)
				} else {
					break
				}
			}
		}
	}()

	wg.Wait()
	close(songCh)
	return songs
}
