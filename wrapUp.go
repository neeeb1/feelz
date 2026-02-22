package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Second)
	return tickMsg{}
}

func updateWrapUp(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tickMsg:
		m.timer--
		if m.timer <= 0 {
			return m, tea.Quit
		}
		return m, tick
	}
	return m, nil
}

func viewWrapUp(m model) string {
	return fmt.Sprintf("Wrapping up in %d seconds... (hit q to quit sooner)", m.timer)
}
