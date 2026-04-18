package main

import (
	"fmt"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	text	string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m model) View() tea.View {
	s := "> " + m.text
	s += fmt.Sprintf("\nPress [enter] to end typing...")
	return tea.NewView(s)
}

func main() {
	now := time.Now()
	fmt.Printf("%v %v %v %v %v:%v:%v\n", now.Weekday().String(), now.Day(), now.Month().String(), now.Year(), now.Hour(), now.Minute(), now.Second())
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		fmt.Printf("There has been an error: %v", err)
		os.Exit(1)
	}
}
