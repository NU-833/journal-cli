package uinp

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	// "github.com/spf13/cobra"
)

type Model struct {
	text	string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case msg.Text != "":
			m.text += msg.Text
		case msg.Code == tea.KeyEnter:
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() tea.View {
	s := fmt.Sprintf("> %v\n", m.text)
	s += fmt.Sprintf("\nPress [enter] to end typing...")
	return tea.NewView(s)
}
