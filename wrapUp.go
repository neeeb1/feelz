package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

//go:embed templates/*
var embedTemplates embed.FS

type notes struct {
	Date    string
	Prompts []journalPrompt
}

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
	var s strings.Builder

	fmt.Fprintf(&s, "Wrapping up in %d seconds... (hit q to quit sooner)\n\n", m.timer)

	for i, p := range m.selected {
		fmt.Fprintf(&s, "%d. %s\n", i+1, p.Prompt)
		s.WriteString(p.FinalText + "\n\n")
	}

	return s.String()
}

func formatMD(m model) (string, error) {
	var s strings.Builder
	note := notes{
		Date:    time.Now().Format(time.RFC1123),
		Prompts: m.selected,
	}

	tmpl, err := template.ParseFS(embedTemplates, "templates/daily.md")
	if err != nil {
		return s.String(), err
	}

	err = tmpl.Execute(&s, note)
	if err != nil {
		return s.String(), err
	}

	return s.String(), nil
}

func writeFile(name, note string) error {
	filePath := filepath.Join("out/", name)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(note)
	if err != nil {
		return err
	}

	file.Sync()

	return nil
}
