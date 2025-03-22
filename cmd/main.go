package main

import (
	"log"

	"github.com/go-telegram/bot"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/0x12th/dev-radar-bot/internal/config"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	_, err = bot.New(cfg.TgBotToken)
	if err != nil {
		log.Printf("[ERROR] failed to create botAPI: %v", err)
		return
	}

	db, err := sqlx.Connect("postgres", cfg.PgDSN)
	if err != nil {
		log.Printf("[ERROR] failed to connect to db: %v", err)
		return
	}
	defer db.Close()

}
