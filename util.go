package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
	NewConfig  bool            `json:"newConfig"`
	OutputPath string          `json:"outputPath"`
	Prompts    []journalPrompt `json:"prompts"`
}

func checkbox(name string, checked bool) string {
	if checked {
		return fmt.Sprintf("[x] %s", name)
	}
	return fmt.Sprintf("[ ] %s", name)
}

func writeError(err string) string {
	return fmt.Sprintf("Error: %s\n\n", err)
}

func loadConfig() (config, error) {
	var conf config

	buf, err := os.ReadFile("config.json")
	if err != nil {
		return conf, err
	}

	json.Unmarshal(buf, &conf)

	return conf, nil
}

func readViperConfig() (config, error) {
	var c config

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return c, err
	}

	viper.SetConfigName(".feelz")
	viper.AddConfigPath(homeDir)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config not found...")
			err = createDefaultViperConfig()
			if err != nil {
				return c, err
			}
		} else {
			return c, err
		}
	}

	if err := viper.Unmarshal(&c); err != nil {
		return c, err
	}

	return c, nil
}

func createDefaultViperConfig() error {
	fmt.Println("Creating new config")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	configPath := filepath.Join(homeDir, ".feelz")

	var prompts = []journalPrompt{
		newJournalPrompt(
			"Today's Thoughts",
			"What are you thinking about today? Any recurring thoughts that you can't get out of your head?").withPlaceholder("How DO magnets work? Miracles?"),
		newJournalPrompt(
			"Graditude is Rad-itude",
			"What are you thankful for today? Try to name 2 or 3 things.").withPlaceholder("No barking from the dog\nNo smog\nMama cooked breakfast with no hog"),
		newJournalPrompt(
			"My Media Diet",
			"What are you consuming lately? Games, movies, music, books... anything!"),
		newJournalPrompt(
			"Today's Top Tasks",
			"Name 3 tasks that are essential for today.\nIf you did these, you've done the bare minimum for a successful day!").withPlaceholder("1. Eat\n2. Brush teeth\n3. ???\n4. Profit!"),
	}
	viper.SetDefault("Prompts", prompts)
	viper.SetDefault("NewConfig", true)
	viper.SetDefault("OutputPath", "out/")

	viper.SetConfigType("json")

	if err := viper.SafeWriteConfigAs(configPath); err != nil {
		return err
	}
	return nil
}
