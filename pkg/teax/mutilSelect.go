package teax

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/get/pkg/teax/stringx"
	"github.com/fzdwx/get/pkg/teax/style"
	"github.com/muesli/termenv"
)

type (
	MultiSelect struct {
		inner *innerSelect
	}

	Option func(ms *MultiSelect)
)

// Show startup MultiSelect
func (ms MultiSelect) Show(text ...string) ([]int, error) {
	WithDefaultText(text...)

	err := ms.inner.Start()
	if err != nil {
		return nil, err
	}

	return ms.inner.Selected(), err
}

// WithDefaultText set default text
func WithDefaultText(text ...string) Option {
	return func(ms *MultiSelect) {
		if len(text) >= 1 {
			ms.inner.defaultText = text[0]
		}
	}
}

// apply options on MultiSelect
func (ms *MultiSelect) apply(ops ...Option) *MultiSelect {
	if len(ops) > 0 {
		for _, option := range ops {
			option(ms)
		}
	}
	return ms
}

/* ============================================================== inner */
type (
	innerSelect struct {
		textStyle   termenv.Style
		choices     []string
		cursor      int
		selected    map[int]struct{}
		defaultText string
	}
)

func newInnerSelect(choices []string) *innerSelect {
	return &innerSelect{
		choices:     choices,
		selected:    make(map[int]struct{}),
		defaultText: "Please select your options",
		textStyle:   style.PrimaryStyle,
	}
}

func (is innerSelect) Selected() []int {
	var selected []int
	for s, _ := range is.selected {
		selected = append(selected, s)
	}
	return selected
}

func (is *innerSelect) Start() error {
	return startUp(is)
}

func (is innerSelect) Init() tea.Cmd {
	return nil
}

func (is *innerSelect) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return is.quit()
		case "up", "k":
			is.moveUp()
		case "down", "j":
			is.moveDown()
		case "enter", " ":
			is.choice()
		}
	}
	return is, nil
}

func (is *innerSelect) View() string {
	msg := stringx.NewFluentSb()

	// The header
	msg.Write(is.textStyle.Styled(is.defaultText)).NextLine()

	// Iterate over our choices
	for i, choice := range is.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if is.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := is.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		msg.Write(fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice))
	}

	// The footer
	msg.Write("\nPress q to quit.\n")

	// Send the UI for rendering
	return msg.String()
}

// moveUp The "up" and "k" keys move the cursor up
func (is *innerSelect) moveUp() {
	if is.cursor > 0 {
		is.cursor--
	}
}

// moveDown The "down" and "j" keys move the cursor down
func (is *innerSelect) moveDown() {
	if is.cursor < len(is.choices)-1 {
		is.cursor++
	}
}

// choice
// The "enter" key and the spacebar (a literal space) toggle
// the selected state for the item that the cursor is pointing at.
func (is *innerSelect) choice() {
	_, ok := is.selected[is.cursor]
	if ok {
		delete(is.selected, is.cursor)
	} else {
		is.selected[is.cursor] = struct{}{}
	}
}

// quit These keys should exit the program.
func (is *innerSelect) quit() (tea.Model, tea.Cmd) {
	return is, tea.Quit
}
