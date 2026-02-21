package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

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

	default:
		return m, tea.Quit
	}
}

func updateSelector(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "up", "k":
			if m.cursorPos > 0 {
				m.cursorPos--
			} else if m.cursorPos == 0 {
				m.cursorPos = len(m.options) - 1
			}

		case "down", "j":
			if m.cursorPos < len(m.options)-1 {
				m.cursorPos++
			} else if m.cursorPos == len(m.options)-1 {
				m.cursorPos = 0
			}

		case "enter", " ":
			if m.cursorPos == len(m.options)-1 {
				m = startSession(m.selected[0])
			} else {
				ok := slices.Contains(m.selected, m.options[m.cursorPos])
				if ok {
					m.selected = slices.DeleteFunc(m.selected, func(E journalPrompt) bool { return E == m.options[m.cursorPos] })
				} else {
					m.selected = append(m.selected, m.options[m.cursorPos])
				}
			}
		}
	}

	return m, nil
}

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

func (m model) View() string {
	switch m.model {
	case TypeSelector:
		return viewSelector(m)

	case TypeSession:
		return viewSession(m)

	default:
		return "whoops"
	}

}

func viewSelector(m model) string {
	var s strings.Builder

	s.WriteString(m.prompt + "\n\n")
	for i, choice := range m.options {

		cursor := " "
		if m.cursorPos == len(m.options)-1 && m.cursorPos == i {
			cursor = "---->"
		} else if m.cursorPos == i {
			cursor = ">"
		}

		checked := false
		if slices.Contains(m.selected, choice) {
			checked = true
		}

		if i == len(m.options)-1 {
			fmt.Fprintf(&s, "\n%s %s", cursor, choice.name)
			if m.cursorPos == len(m.options)-1 {
				s.WriteString("!")
			}
			s.WriteString("\n")
		} else {
			fmt.Fprintf(&s, "%s %s\n", cursor, checkbox(choice.name, checked))
		}
	}

	s.WriteString("\nPress q to quit.\n")
	return s.String()
}

func viewSession(m model) string {
	var s strings.Builder

	s.WriteString(m.prompt + "\n\n")
	s.WriteString(m.textarea.View())
	s.WriteString("\nPress q to quit.\n")
	return s.String()
}

func checkbox(name string, checked bool) string {
	if checked {
		return fmt.Sprintf("[x] %s", name)
	}
	return fmt.Sprintf("[ ] %s", name)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("An error occured: %v", err)
		os.Exit(1)
	}
}
