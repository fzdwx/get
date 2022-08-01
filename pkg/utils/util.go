package utils

import (
	"io"
	"net/http"
)

var debug = false

// SetDebug set env to debug
func SetDebug() {
	debug = true
}

// IsDebug Current env is debug?
func IsDebug() bool {
	return debug
}

// ReadBody request body and close it.
func ReadBody(readerCloser io.ReadCloser) []byte {
	defer readerCloser.Close()
	bytes, err := io.ReadAll(readerCloser)
	if err != nil {
		return []byte{}
	}
	return bytes
}

// MappingArtName mapping art name,default "未知"
func MappingArtName(name string) string {
	if name == "" {
		return "未知"
	}
	return name
}

// AdapterScreenTruncate truncate string adapter screen width
func AdapterScreenTruncate(s string) string {
	width := GetTermWidth()

	if width != 0 {
		width -= 5
	}

	return Truncate(s, width)
}

// GetSize request url get headers[content-length]
func GetSize(url string) int64 {
	resp, err := http.Get(url)
	if err != nil {
		return 0
	}

	return resp.ContentLength
}
