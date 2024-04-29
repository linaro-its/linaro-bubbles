package main

import (
	"fmt"
	"os"

	"github.com/linaro-its/linaro-bubbles/picker"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	linpick picker.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	m.linpick, cmd = m.linpick.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	return m.linpick.View()
}

func main() {
	items := []picker.Item{
		{Key: "Alpha", Value: "1"},
		{Key: "Bravo", Value: "2"},
		{Key: "Charlie", Value: "3"},
		{Key: "Delta", Value: "4"},
	}
	input := Model{linpick: picker.New(items)}
	input.linpick.SetCursor("3")
	p := tea.NewProgram(input)
	m, err := p.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if m, ok := m.(Model); ok {
		cursor := m.linpick.Cursor()
		key := m.linpick.Items()[cursor].Key
		value := m.linpick.Items()[cursor].Value
		fmt.Println(cursor, key, value)
	}
}
