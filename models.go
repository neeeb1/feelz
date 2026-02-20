package main

type model struct {
	prompt         string
	journalPrompts []journalPrompt
	cursor         int
	selected       map[int]struct{}
}

type journalPrompt struct {
	name   string
	prompt string
}

func newJournalPrompt(name, prompt string) journalPrompt {
	return journalPrompt{
		name:   name,
		prompt: prompt,
	}
}

func initialModel() model {
	return model{
		prompt: "What do you want to write about today? Choose a few prompts to get started.",
		journalPrompts: []journalPrompt{
			newJournalPrompt("Today's Thoughts", "What are you thinking about today? Any recurring thoughts that you can't get out of your head?"),
			newJournalPrompt("Graditude is Rad-itude", "What are you thankful for today?"),
			newJournalPrompt("My Media Diet", "What are you consuming lately? Games, movies, music, books... anything!"),
		},

		selected: make(map[int]struct{}),
	}
}

func journalSession(p journalPrompt) model {
	modelPrompt := p.prompt
	return model{
		prompt: modelPrompt,
	}
}
