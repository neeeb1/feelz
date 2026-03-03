package main

import (
	"encoding/json"
	"fmt"
	"os"

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

func viperConfig() (config, error) {
	var c config

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return c, err
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		return c, err
	}

	return c, nil
}
