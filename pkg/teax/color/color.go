package color

import (
	"github.com/muesli/termenv"
	"strconv"
)

var (
	Color = termenv.ColorProfile()

	DefaultFg   = NewColor(39)
	DefaultBg   = NewColor(49)
	FgLightCyan = NewColor(96)
)

func NewColor(i int) termenv.Color {
	return Color.Color(strconv.Itoa(i))
}
