package main

import (
	"log"

	"github.com/go-telegram/bot"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/0x12th/dev-radar-bot/internal/config"
)


func main() {
	_, err := bot.New(config.GetConfig().TgBotToken)
	if err != nil {
		log.Printf("[ERROR] failed to create botAPI: %v", err)
		return
	}

	db, err := sqlx.Connect("postgres", config.GetConfig().PgDSN)
	if err != nil {
		log.Printf("[ERROR] failed to connect to db: %v", err)
		return
	}
	defer db.Close()

}