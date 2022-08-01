package utils

import (
	"fmt"
	"testing"
)

func TestEncodeToUrl(t *testing.T) {
	fmt.Println(EncodeToUrl("你好"))
}

func TestTruncate(t *testing.T) {
	s := " Stay (cover: Dame Dame|Majes|Jordan Rys|Charlton Howard|Justin Bieber|Magnus Høiberg|Charlie Puth|Omer Fedi|Blake Slatkin|Michael MuleIsaac De Boni|Subhaan Rahmaan)(685.84KB)"
	fmt.Println(Truncate(s, 20))
	fmt.Println(Truncate("123", 20))
}
