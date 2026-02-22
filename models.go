package main

import (
	"github.com/charmbracelet/bubbles/textarea"
)

type ModelType int

const (
	TypeSelector ModelType = iota
	TypeSession
)

type model struct {
	prompt    string
	model     ModelType
	cursorPos int
	selected  []journalPrompt

	// Fields for selector type
	options []journalPrompt

	// Fields for session type
	textarea      textarea.Model
	currentPrompt int
}

type journalPrompt struct {
	name            string
	prompt          string
	placeholderText string
}

func newJournalPrompt(name, prompt string) journalPrompt {
	return journalPrompt{
		name:   name,
		prompt: prompt,
	}
}

func (p journalPrompt) withPlaceholder(ph string) journalPrompt {
	p.placeholderText = ph
	return p
}

func initialModel() model {
	prompts := []journalPrompt{
		newJournalPrompt("Today's Thoughts", "What are you thinking about today? Any recurring thoughts that you can't get out of your head?"),
		newJournalPrompt("Graditude is Rad-itude", "What are you thankful for today?"),
		newJournalPrompt("My Media Diet", "What are you consuming lately? Games, movies, music, books... anything!"),

		newJournalPrompt("Get started", ""),
	}

	return model{
		prompt:   "What do you want to write about today? Choose a few prompts to get started.",
		options:  prompts,
		selected: make([]journalPrompt, 0, len(prompts)),
		model:    TypeSelector,
	}
}

func startSession(p journalPrompt) model {
	modelPrompt := p.prompt

	input := textarea.New()
	input.Placeholder = p.placeholderText
	input.Focus()

	return model{
		prompt:   modelPrompt,
		model:    TypeSession,
		options:  make([]journalPrompt, 0),
		selected: make([]journalPrompt, 0),
		textarea: input,
	}
}
