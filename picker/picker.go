package picker

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	Prompt string
}

func (m Model) View() string {
	// Render the currently selected item
	return ""
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}
