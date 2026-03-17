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

	msg, err := client.SendPhoto(telegram.SendPhotoRequest{
		ChatID:  "@move_it_move",
		Photo:   "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQqad93DYG24AlwsrIQtEnhwsnFrxiTuj6P0Q&s",
		Caption: "This photo is sent via BOT API",
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
