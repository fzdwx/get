package msc

import (
	"fmt"
	"github.com/fzdwx/get/pkg/utils"
	"github.com/pterm/pterm"
	"github.com/rotisserie/eris"
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
		// e.g: .mp3
		EncodeType string `json:"encodeType"`
	}

	Artists struct {
		name string
	}
)

// Platform the music service provider
type Platform int

const (
	NetEasyP Platform = 1
	KuWoP    Platform = 2
)

func (s Songs) Prompt(i int) string {
	return utils.AdapterScreenTruncate(fmt.Sprintf("%d. %s(%s) - %s", i, s.Name, utils.FormatBytes(s.Size), utils.MappingArtName(s.Artists[0].name)))
}

// Collect songs
// async call SongsMapper#mapper() converted to Songs
func Collect(mappers []SongsMapper) []Songs {
	songCh := make(chan Songs)
	wg := sync.WaitGroup{}
	wg.Add(len(mappers))

	var songs []Songs
	go func() {
		for i := range mappers {
			go func(mapper SongsMapper) {
				mscSongs, err := mapper.mapper()
				if err != nil {
					pterm.Error.Printfln(eris.ToString(err, true))
					wg.Done()
				} else {
					songCh <- *mscSongs
				}

			}(mappers[i])
		}
	}()

	go func() {
		for {
			select {
			case s, ok := <-songCh:
				if ok {
					songs = append(songs, s)
					wg.Done()
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
