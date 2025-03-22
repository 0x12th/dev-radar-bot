package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	PgDSN      string
	TgBotToken string
	OpenAIKey  string
}

var (
	cfg  *Config
	once sync.Once
)

func GetConfig() (*Config, error) {
	once.Do(func() {
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Warning: unable to read config file: %v", err)
		}

		cfg = &Config{
			PgDSN:      viper.GetString("PgDSN"),
			TgBotToken: viper.GetString("TgBotToken"),
			OpenAIKey:  viper.GetString("OpenAIKey"),
		}
	})

	if cfg == nil {
		return nil, fmt.Errorf("[ERROR] missing required configuration fields")
	}

	return cfg, nil
}
