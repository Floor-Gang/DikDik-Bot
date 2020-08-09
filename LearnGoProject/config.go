package main

import (
	util "github.com/Floor-Gang/utilpkg/config"
	"log"
)

type DikDikConfig struct {
	Token        string            `yaml:"bot_token"`
	Prefix       string            `yaml:"bot_prefix"`
	CommandTitle string            `yaml:"command_title"`
	Commands     []string          `yaml:"commands"`
	Channels     map[string]string `yaml:"stored_channels"`
}
type channel struct{
	Topic	string
}

const configPath = "./dikdik-config.yml"

func GetConfig() DikDikConfig {
	var prefix = "/"
	defaultConfig := DikDikConfig{
		Prefix:       prefix,
		CommandTitle: "Commands",
		Commands: []string{
			prefix + "+say channelName message to send to channel",
			prefix + "-say",
			prefix + "jokeHere",
			prefix + "jokeThere MentionChannel",
			prefix + "factsHere",
			prefix + "factsThere MentionChannel",
			prefix + "help",
		},
	}
	err := util.GetConfig(configPath, &defaultConfig)

	if err != nil {
		log.Fatalln(err)
	}
	return defaultConfig
}

func (config DikDikConfig) Save() {
	util.Save(configPath, &config)
}