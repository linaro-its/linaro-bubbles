package picker

import tea "github.com/charmbracelet/bubbletea"

type Item struct {
	Key   string
	Value string
}

type Model struct {
	cursor int
	items  []Item
}

func (m Model) View() string {
	if len(m.items) == 0 {
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

func (m Model) Items() []Item {
	return m.items
}

func (m *Model) SetItems(i []Item) {
	m.items = i
}

func New(items []Item) Model {
	m := Model{
		cursor: 0,
		items:  items,
	}
	return m
}
