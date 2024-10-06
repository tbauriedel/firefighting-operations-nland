package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/tbauriedel/firefighting-operations-nland/internal/config"
	"github.com/tbauriedel/firefighting-operations-nland/internal/operations"
	"github.com/tbauriedel/firefighting-operations-nland/internal/scraper"
	"github.com/tbauriedel/firefighting-operations-nland/internal/telegram"
	"log"
	"os"
	"time"
)

const product = "firefighting-operations-nland"

const readme = "Scrapes firefighting operations from kfv-online.de and sends them to telegram via api"

var (
	createConfig      bool
	lastSentOperation operations.Operation
)

func init() {
	flag.BoolVar(&createConfig, "generate-config", false, fmt.Sprintf("Create example config file. Will be saved into %s", config.DefaultConfigDir))

	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "**%s**\n%s\n\n", product, readme)

		_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])

		flag.PrintDefaults()
	}

	flag.Parse()

	var err error

	if createConfig {
		err = config.CreateDefaultConfigFile()
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	config.Config, err = config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	tBot, err := telegram.NewBotInstance(config.Config.TelegramBotID)
	if err != nil {
		log.Fatal(err)
	}

	for {
		s := scraper.New()
		s.RegisterOperations()
		handleOperations(s, tBot)

		time.Sleep(config.Config.ScraperInterval)
	}
}

func handleOperations(s *scraper.Scraper, t telegram.Bot) {
	err := s.Collector.Visit("https://www.kfv-online.de/home/einsaetze")
	if err != nil {
		log.Printf("error while vistiting url: %s", err)
	}

	if s.RC != 200 || s.Failure {
		return
	}

	if s.RC == 200 && !s.Failure {
		lastFoundOperation := s.Operations[0]

		if lastSentOperation != lastFoundOperation {
			log.Print("new operation found!")
			sendErr := t.Send(config.Config.TelegramChatID, buildMessage(s.Operations[0]))
			if sendErr != nil {
				log.Fatal(sendErr)
			}
			lastSentOperation = lastFoundOperation
		}
	}
}

func buildMessage(o operations.Operation) (message string) {
	message = "\U0001F692 " + o.Units + "\n" + "\U0001F4A5 " + o.Report + "\n" + "\U0001F4CD " + o.Location + " (" + o.District + ")\n" + "\U0001F551 " + o.Time
	return
}
