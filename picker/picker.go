package picker

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Item struct {
	Key   string
	Value string
}

type Model struct {
	cursor      int
	focus       bool
	empty       string
	items       []Item
	Prompt      string
	TextStyle   lipgloss.Style
	PromptStyle lipgloss.Style
}

func (m Model) View() string {
	var value string
	if len(m.items) == 0 {
		if m.empty != "" {
			value = m.empty
		} else {
			value = "no items"
		}
	} else {
		value = m.items[m.cursor].Key
	}

	return m.PromptStyle.Render(m.Prompt) + m.TextStyle.Render(value)
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if !m.focus {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left":
			if m.cursor > 0 {
				m.cursor--
			}
		case "right":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		}
	}

	return m, nil
}

func (m Model) Cursor() int {
	return m.cursor
}

func (m *Model) SetCursor(value string) error {
	// Set the cursor by finding the specified value
	// in the list.
	for i := 0; i < len(m.items); i++ {
		if m.items[i].Value == value {
			m.cursor = i
			return nil
		}
	}

	return fmt.Errorf("failed to match value '%s'", value)
}

func (m *Model) Focus() {
	m.focus = true
}

func (m *Model) Blur() {
	m.focus = false
}

func (m Model) Items() []Item {
	return m.items
}

func (m *Model) SetItems(i []Item) {
	m.items = i
}

func (m *Model) SetEmpty(i string) {
	m.empty = i
}

func New(items []Item) Model {
	m := Model{
		cursor: 0,
		focus:  false,
		items:  items,
	}
	return m
}
