package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

// Config структура для конфигурации
type Config struct {
	PgDSN      string
	TgBotToken string
	OpenAIKey  string
}

var (
	cfg  Config
	once sync.Once
)

func GetConfig() Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		cfg.PgDSN = os.Getenv("PgDSN")
		cfg.TgBotToken = os.Getenv("TgBotToken")
		cfg.OpenAIKey = os.Getenv("OpenAIKey")

		if cfg.PgDSN == "" || cfg.TgBotToken == "" || cfg.OpenAIKey == "" {
			log.Fatal("[ERROR] missing required configuration fields")
		}

	})
	return cfg
}