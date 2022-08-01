package utils

import "github.com/pterm/pterm"

// GetTermWidth get term width
func GetTermWidth() int {
	w, _, err := pterm.GetTerminalSize()
	if err != nil {
		return 0
	}

	return w
}
