package msc

import (
	"encoding/json"
	"fmt"
	"github.com/fzdwx/get/pkg/utils"
	"github.com/rotisserie/eris"
	"net/http"
	"os"
)

type (
	netEasy struct {
		name    string
		pageNum int
	}

	netEasySearchResponse struct {
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
	netEasySearchUrl = "https://music.163.com/api/search/get/web?s=%s&type=1&limit=20&offset=%d"
	netEasyDetailUrl = "https://music.163.com/api/song/enhance/player/url?id=%d&ids=[%d]&br=3200000"
)

func newNetEasy(name string) Request {
	return &netEasy{
		pageNum: 1,
		name:    utils.EncodeToUrl(name),
	}
}

func (n *netEasy) Execute() ([]Songs, int, error) {
	url := n.url()
	if utils.IsDebug() {
		fmt.Fprintf(os.Stderr, "search url:%s\n", url)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, 0, eris.Wrapf(err, "netEasy search fail. url: %s", url)
	}

	var result netEasySearchResponse
	body := utils.ReadBody(resp.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, 0, eris.Wrapf(err, "format netEasy search json fail. %s", string(body))
	}

	var mappers []SongsMapper
	for _, song := range result.Result.Songs {
		mappers = append(mappers, SongsMapper(song))
	}

	return Collect(mappers), result.Result.SongCount, nil
}

func (n *netEasy) url() string {
	offset := (n.pageNum - 1) * 20
	return fmt.Sprintf(netEasySearchUrl, n.name, offset)
}

func (n *netEasy) prevPage() {
	if n.pageNum >= 1 {
		n.pageNum = n.pageNum - 1
	}
}

func (n *netEasy) nextPage() {
	n.pageNum = n.pageNum + 1
}

func (ns netEasySong) name() string {
	return ns.Name
}

func (ns netEasySong) mapper() (*Songs, error) {
	url := fmt.Sprintf(netEasyDetailUrl, ns.Id, ns.Id)
	if utils.IsDebug() {
		fmt.Fprintf(os.Stderr, "data url:%s\n", url)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, eris.Wrapf(err, "get netEasy music fail. url: %s", url)
	}

	var result netDataResponse
	body := utils.ReadBody(resp.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, eris.Wrapf(err, "format netEasy data json fail. %s", string(body))
	}

	if len(result.Data) < 1 {
		return nil, eris.New(fmt.Sprintf("not found netEasy songs. url: %s", url))
	}

	data := result.Data[0]
	return &Songs{
		Name:        ns.Name,
		Artists:     ns.Artists,
		DownloadUrl: data.Url,
		Size:        data.Size,
		EncodeType:  "." + data.EncodeType,
	}, nil
}
