package msc

import (
	"fmt"
	"testing"
)

func Test_netEasy(t *testing.T) {
	songs, count, _ := newNetEasy("稻香").Execute()
	fmt.Println(count)
	fmt.Println(songs)
}
