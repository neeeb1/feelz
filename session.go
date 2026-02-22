package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func updateSession(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case "ctrl+n":
			m.selected[m.currentPrompt].FinalText = m.textarea.Value()
			m.currentPrompt += 1

			if m.currentPrompt <= len(m.selected)-1 {
				m = startSession(m.selected[m.currentPrompt], m)
			} else {
				m = wrapUp(m)
				cmds = append(cmds, tick)
			}
		case "ctrl+c", "q":
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

	fmt.Fprintf(&s, "%d. %s\n\n", m.currentPrompt+1, m.selected[m.currentPrompt].Name)
	s.WriteString(m.prompt + "\n\n")
	s.WriteString(m.textarea.View())
	s.WriteString("\nPress q to quit.\n")
	return s.String()
}
