package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if msg, ok := msg.(tea.KeyMsg); ok {
		s := msg.String()
		if s == "ctrl+c" || s == "q" {
			return m, tea.Quit
		}
	}

	switch m.model {

	case TypeSelector:
		return updateSelector(msg, m)

	case TypeSession:
		return updateSession(msg, m)

	case TypeWrapUp:
		return updateWrapUp(msg, m)

	default:
		return m, tea.Quit
	}
}

func (m model) View() string {
	switch m.model {
	case TypeSelector:
		return viewSelector(m)

	case TypeSession:
		return viewSession(m)

	case TypeWrapUp:
		return viewWrapUp(m)

	default:
		return "whoops"
	}

}

func main() {
	conf, err := loadConfig()
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	m := initialModel(conf)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("An error occured: %v", err)
		os.Exit(1)
	}
}
