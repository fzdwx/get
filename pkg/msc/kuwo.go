package msc

import (
	"encoding/json"
	"fmt"
	"github.com/fzdwx/get/pkg/utils"
	"github.com/rotisserie/eris"
	"net/http"
	"os"
	"path/filepath"
)

type (
	kuWo struct {
		name    string
		pageNum int
	}

	kuWoSearchResponse struct {
		Total   string      `json:"Total"`
		AbsList []absEntity `json:"abslist"`
	}

	absEntity struct {
		// songs name
		Name string `json:"Name"`
		// songs id
		Id string `json:"DC_TARGETID"`
		// songs author
		Artist string `json:"ARTIST"`
	}

	kuWoDataResponse struct {
		Msg  string   `json:"msg"`
		Data kuWoData `json:"data"`
	}

	kuWoData struct {
		Url string `json:"url"`
	}
)

const (
	kuWoSearchUrl = "https://search.kuwo.cn/r.s?client=kt&all=%s&pn=%d&rn=20&vipver=1&ft=music&encoding=utf8&rformat=json&mobi=1"
	kuWoDetailUrl = `https://www.kuwo.cn/api/v1/www/music/playUrl?mid=%s&type=1`
)

func newKuWo(name string) Request {
	return &kuWo{
		name:    utils.EncodeToUrl(name),
		pageNum: 1,
	}
}

func (k *kuWo) Execute() ([]Songs, int, error) {
	url := k.url()
	if utils.IsDebug() {
		fmt.Fprintf(os.Stderr, "search url:%s\n", url)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, 0, eris.Wrapf(err, "kuWo search fail. url: %s", url)
	}

	var result kuWoSearchResponse
	body := utils.ReadBody(resp.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, 0, eris.Wrapf(err, "format kuWo search json fail. %s", string(body))
	}

	var mappers []SongsMapper
	for _, abs := range result.AbsList {
		mappers = append(mappers, SongsMapper(abs))
	}

	return Collect(mappers), utils.ToInt(result.Total), nil
}

func (k *kuWo) url() string {
	offset := k.pageNum - 1
	return fmt.Sprintf(kuWoSearchUrl, k.name, offset)
}

func (k *kuWo) prevPage() {
	if k.pageNum >= 1 {
		k.pageNum = k.pageNum - 1
	}
}

func (k *kuWo) nextPage() {
	k.pageNum = k.pageNum + 1
}

func (a absEntity) name() string {
	return a.Name
}

func (a absEntity) mapper() (*Songs, error) {
	url := fmt.Sprintf(kuWoDetailUrl, a.Id)
	if utils.IsDebug() {
		fmt.Fprintf(os.Stderr, "data url:%s\n", url)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, eris.Wrapf(err, "get kuWo music fail. url: %s", url)
	}

	var result kuWoDataResponse
	body := utils.ReadBody(resp.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, eris.Wrapf(err, "format kuWo data json fail. %s", string(body))
	}

	if !(result.Msg == "success") {
		return nil, eris.New(fmt.Sprintf("not found kuWo songs. url: %s", url))
	}

	return &Songs{
		Name: a.Name,
		Artists: []Artists{
			{name: a.Artist},
		},
		DownloadUrl: result.Data.Url,
		// todo 能异步吗？
		Size:       utils.GetSize(result.Data.Url),
		EncodeType: filepath.Ext("https://sr-sycdn.kuwo.cn/e34dbb6596ab939508d3b20745a8a197/62e66490/resource/n1/43/63/3458986373.mp3"),
	}, nil
}
