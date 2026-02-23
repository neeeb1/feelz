package main

import (
	"github.com/charmbracelet/bubbles/textarea"
)

type ModelType int

const (
	TypeSelector ModelType = iota
	TypeSession
	TypeWrapUp
)

type model struct {
	prompt    string
	model     ModelType
	cursorPos int
	selected  []journalPrompt
	err       string

	// Fields for selector type
	options []journalPrompt

	// Fields for session type
	textarea      textarea.Model
	currentPrompt int

	// Fields for wrapUp type
	timer int
	final string
}

type journalPrompt struct {
	Name            string
	Prompt          string
	PlaceholderText string
	FinalText       string
}

func newJournalPrompt(name, prompt string) journalPrompt {
	return journalPrompt{
		Name:   name,
		Prompt: prompt,
	}
}

func (p journalPrompt) withPlaceholder(ph string) journalPrompt {
	p.PlaceholderText = ph
	return p
}

func initialModel(conf config) model {
	var prompts []journalPrompt
	prompts = conf.PromptTemplates
	prompts = append(prompts, newJournalPrompt("Get Started", ""))

	return model{
		prompt:   "What do you want to write about today? Choose a few prompts to get started.",
		options:  prompts,
		selected: make([]journalPrompt, 0, len(prompts)),
		model:    TypeSelector,
	}
}

func startSession(p journalPrompt, prevModel model) model {
	modelPrompt := p.Prompt

	input := textarea.New()
	input.Placeholder = p.PlaceholderText
	input.Focus()

	return model{
		prompt:        modelPrompt,
		model:         TypeSession,
		textarea:      input,
		currentPrompt: prevModel.currentPrompt,
		selected:      prevModel.selected,
	}
}

func wrapUp(prevModel model) model {
	newModel := model{
		selected: prevModel.selected,
		model:    TypeWrapUp,
		timer:    5,
	}

	note, err := formatMD(prevModel)
	if err != nil {
		newModel.err = err.Error()
	}

	if err = writeFile("test", note); err != nil {
		newModel.err = err.Error()
	}

	return newModel
}
