package msc

import (
	"fmt"
	"github.com/fzdwx/get/pkg"
)

type (
	Request interface {
		// execute request get songs
		execute() ([]Songs, int, error)
		// generate request url
		url() string
		// next page
		nextPage()
		// prev page
		prevPage()
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

func (s Songs) Prompt(i int) string {
	return fmt.Sprintf("%d. %s(%s) - %s", i, s.Name, pkg.FormatBytes(s.Size), pkg.MappingArtName(s.Artists[0].name))
}
