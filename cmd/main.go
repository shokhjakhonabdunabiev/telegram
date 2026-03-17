package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/shokhjakhonabdunabiev/telegram"
)

type Config struct {
	botApiToken string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	botApiToken := os.Getenv("BOT_API_TOKEN")

	cfg := Config{
		botApiToken: botApiToken,
	}

	client := telegram.NewClient(cfg.botApiToken, 10*time.Second)

	msg, err := client.SendMessage(telegram.SendMessageRequest{
		ChatID:    "@move_it_move",
		Text:      "Wrong parse_mode",
		ParseMode: "HTML",
	})

	if err != nil {
		fmt.Println(err)
	} else {
		print(msg)
	}
}

func print[T any](data T) {
	b, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(b))
}
