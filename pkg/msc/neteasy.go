package msc

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/fzdwx/get/pkg/utils"
	"github.com/pterm/pterm"
)

type (
	netEasy struct {
		name    string
		pageNum int
	}

	netEasyResponse struct {
		Result netEasyResult `json:"result"`
	}

	netEasyResult struct {
		SongCount int           `json:"songCount"`
		Songs     []netEasySong `json:"songs"`
	}

	netEasySong struct {
		Id      int       `json:"id"`
		Name    string    `json:"Name"`
		Artists []Artists `json:"Artists"`
	}

	netDataResponse struct {
		Data []netData `json:"data"`
	}

	netData struct {
		Url        string `json:"url"`
		Size       int64  `json:"Size"`
		EncodeType string `json:"EncodeType"`
	}
)

const (
	netEasyUrl       = "https://music.163.com/api/search/get/web?s=%s&type=1&limit=20&offset=%d"
	netEasyDetailUrl = "https://music.163.com/api/song/enhance/player/url?id=%d&ids=[%d]&br=3200000"
)

func newNetEasy(name string) *netEasy {
	return &netEasy{
		pageNum: 1,
		name:    utils.EncodeToUrl(name),
	}
}

func (n *netEasy) execute() ([]Songs, int, error) {
	resp, err := http.Get(n.url())
	if err != nil {
		return nil, 0, err
	}

	var result netEasyResponse
	body := utils.ReadBody(resp.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, 0, err
	}

	songCh := make(chan Songs)
	wg := sync.WaitGroup{}
	wg.Add(len(result.Result.Songs))

	var songs []Songs
	go func() {
		for i := range result.Result.Songs {
			go func(song netEasySong) {
				mscSongs, err := song.execute()
				if err != nil {
					wg.Done()
					pterm.Error.Printfln("download %s fail", song.Name)
				}

				songCh <- *mscSongs
			}(result.Result.Songs[i])
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
	return songs, result.Result.SongCount, nil
}

func (n *netEasy) url() string {
	offset := (n.pageNum - 1) * 20
	return fmt.Sprintf(netEasyUrl, n.name, offset)
}

func (n *netEasy) prevPage() {
	if n.pageNum >= 1 {
		n.pageNum = n.pageNum - 1
	}
}

func (n *netEasy) nextPage() {
	n.pageNum = n.pageNum + 1
}

func (ns netEasySong) execute() (*Songs, error) {
	resp, err := http.Get(fmt.Sprintf(netEasyDetailUrl, ns.Id, ns.Id))
	if err != nil {
		return nil, err
	}

	var result netDataResponse
	body := utils.ReadBody(resp.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Data) < 1 {
		return nil, errors.New("not found songs")
	}

	data := result.Data[0]
	return &Songs{
		Name:        ns.Name,
		Artists:     ns.Artists,
		DownloadUrl: data.Url,
		Size:        data.Size,
		EncodeType:  data.EncodeType,
	}, nil
}
