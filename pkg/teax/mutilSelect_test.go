package teax

import (
	"fmt"
	"testing"
)

func TestMultiSelect_Show(t *testing.T) {
	selected := NewMultiSelect([]string{"Buy carrots", "Buy celery", "Buy kohlrabi"})
	s, err := selected.Show()
	if err != nil {
		return
	}
	fmt.Println(s)
}
