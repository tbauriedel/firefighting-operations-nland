package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
	"path"
	"time"
)

const DefaultConfigDir = "/etc/firefighting-operations-nland"
const DefaultUser = "firefighting-operations-nland"
const DefaultGroup = "firefighting-operations-nland"

type c struct {
	TelegramBotID   string        `yaml:"telegram_bot_id"`
	TelegramChatID  int64         `yaml:"telegram_chat_id"`
	ScraperInterval time.Duration `yaml:"scraper_interval"`
}

var Config c

func GetConfigDefaults() c {
	return c{
		TelegramBotID:   "",
		TelegramChatID:  0,
		ScraperInterval: time.Second * 10,
	}
}

func ReadConfig() (c, error) {

	fileData, err := os.ReadFile(path.Join(DefaultConfigDir, "config.yaml"))
	if err != nil {
		return c{}, fmt.Errorf("failed to read config.yaml: %w", err)
	}

	conf := c{}

	err = yaml.Unmarshal(fileData, &conf)
	if err != nil {
		return c{}, fmt.Errorf("failed to unmarshal config.yaml: %w", err)
	}

	log.Print("Read config file successfully")

	return conf, nil
}

func CreateDefaultConfigFile() error {
	log.Print("Creating default config file")

	file, err := os.Create(path.Join(DefaultConfigDir, "config.yaml"))
	if err != nil {
		return fmt.Errorf("could not create default config file: %w", err)
	}

	defer file.Close()

	defaults := GetConfigDefaults()

	yamlData, err := yaml.Marshal(&defaults)
	if err != nil {
		return fmt.Errorf("could not marshal default config: %w", err)
	}

	_, err = io.WriteString(file, string(yamlData))
	if err != nil {
		return fmt.Errorf("error writing default config: %w", err)
	}

	log.Print("Finished creating default config file")

	return nil
}
