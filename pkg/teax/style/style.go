package style

import (
	"github.com/fzdwx/get/pkg/teax/color"
	"github.com/muesli/termenv"
)

func New() termenv.Style {
	return termenv.String()
}

var (
	DefaultStyle = New().Foreground(color.DefaultFg).Background(color.DefaultBg)
	PrimaryStyle = New().Foreground(color.DefaultFg)
)
