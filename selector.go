package main

import (
	"fmt"
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

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
				if len(m.selected) != 0 {
					m.currentPrompt = 0
					m = startSession(m.selected[m.currentPrompt], m)
				} else {
					m.err = "Please select an option first!"
				}
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

func viewSelector(m model) string {
	var s strings.Builder

	s.WriteString(m.prompt + "\n\n")
	if m.err != "" {
		s.WriteString(writeError(m.err))
	}
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
