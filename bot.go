package telegram

import (
	"encoding/json"
)

func (bot *Bot) GetMe() (*User, error) {
	res, err := bot.get("getMe", nil)
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(res.Result, &user)

	return &user, err
}

type GetChatRequest struct {
	ChatID string `json:"chat_id"`
}

func (bot *Bot) GetChat(params GetChatRequest) (*ChatFullInfo, error) {
	res, err := bot.get("getChat", params)
	if err != nil {
		return nil, err
	}

	var chatInfo ChatFullInfo
	err = json.Unmarshal(res.Result, &chatInfo)

	return &chatInfo, err
}

type ParseMode string

const (
	HTML       ParseMode = "HTML"
	MarkdownV2 ParseMode = "MarkdownV2"
)

// TODO: reply_markup can be InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
type SendMessageRequest struct {
	BusinessConnectionID    string                   `json:"business_connection_id,omitempty"`
	ChatID                  string                   `json:"chat_id"`
	MessageThreadID         int                      `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                      `json:"direct_messages_topic_id,omitempty"`
	Text                    string                   `json:"text"`
	ParseMode               ParseMode                `json:"parse_mode,omitempty"`
	Entities                []*MessageEntity         `json:"entities,omitempty"`
	LinkPreviewOptions      *LinkPreviewOptions      `json:"link_preview_options,omitempty"`
	DisableNotification     bool                     `json:"disable_notification,omitempty"`
	ProtectContent          bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                     `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                   `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                      `json:"reply_markup,omitempty"`
}

func (bot *Bot) SendMessage(body SendMessageRequest) (*Message, error) {
	res, err := bot.post("sendMessage", body)
	if err != nil {
		return nil, err
	}

	var msg Message
	err = json.Unmarshal(res.Result, &msg)

	return &msg, err
}

type SendPhotoRequest struct {
	ChatID  string `json:"chat_id"`
	Photo   string `json:"photo"`
	Caption string `json:"caption,omitempty"`
}

func (bot *Bot) SendPhoto(body SendPhotoRequest) (*Message, error) {
	res, err := bot.post("sendPhoto", body)
	if err != nil {
		return nil, err
	}

	var msg Message
	err = json.Unmarshal(res.Result, &msg)

	return &msg, err
}
