package pkg

import (
	"fmt"
	"io"
	"net/url"
	"strings"
)

// EncodeToUrl encode val to urlEncode
func EncodeToUrl(val string) string {
	return url.QueryEscape(val)
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

func MappingArtName(name string) string {
	if name == "" {
		return "未知"
	}
	return name
}

func NormalizeFileName(name string) string {
	return strings.ReplaceAll(name, ":", "")
}

const (
	BYTE = 1.0 << (10 * iota)
	KILOBYTE
	MEGABYTE
	GIGABYTE
	TERABYTE
)

func FormatBytes(bytes int64) string {
	unit := ""
	value := float32(bytes)

	switch {

	case bytes >= TERABYTE:
		unit = "TB"
		value = value / TERABYTE
	case bytes >= GIGABYTE:
		unit = "GB"
		value = value / GIGABYTE
	case bytes >= MEGABYTE:
		unit = "MB"
		value = value / MEGABYTE
	case bytes >= KILOBYTE:
		unit = "KB"
		value = value / KILOBYTE
	case bytes == 0:
		return "0"

	}

	result := fmt.Sprintf("%.2f", value)
	result = strings.TrimSuffix(result, ".00")
	return fmt.Sprintf("%s%s", result, unit)
}
