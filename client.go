package telegram

import (
	"net/http"
	"time"
)

const (
	baseURL = "https://api.telegram.org/bot"
)

type Bot struct {
	token      string
	baseURL    string
	httpClient *http.Client
}

func NewClient(token string, timeout time.Duration) *Bot {
	return &Bot{
		token:   token,
		baseURL: baseURL + token,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}
