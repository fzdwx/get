package msc

import (
	"testing"
)

func Test_kuWo(t *testing.T) {
	_, _, err := newKuWo("周杰伦").Execute()
	if err != nil {
		panic(err)
	}
}

func Test_URL(t *testing.T) {
}
