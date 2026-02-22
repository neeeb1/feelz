package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func updateSession(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

		m.textarea, cmd = m.textarea.Update(msg)
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)

}

func viewSession(m model) string {
	var s strings.Builder

	s.WriteString(m.prompt + "\n\n")
	s.WriteString(m.textarea.View())
	s.WriteString("\nPress q to quit.\n")
	return s.String()
}
