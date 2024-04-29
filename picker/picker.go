package picker

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Item struct {
	Key   string
	Value string
}

type Model struct {
	cursor int
	empty  string
	items  []Item
}

func (m Model) View() string {
	if len(m.items) == 0 {
		if m.empty != "" {
			return m.empty
		}
		return "no items"
	}
	// Render the currently selected item
	return m.items[m.cursor].Key
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
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
		items:  items,
	}
	return m
}
