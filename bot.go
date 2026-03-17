package telegram

import (
	"encoding/json"
	"fmt"
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
	Markdown   ParseMode = "Markdown"
	MarkdownV2 ParseMode = "MarkdownV2"
)

func (p ParseMode) isValid() bool {
	switch p {
	case HTML, Markdown, MarkdownV2:
		return true
	}
	return false
}

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
	ReplyMarkup             ReplyMarkup              `json:"reply_markup,omitempty"`
} 

func (bot *Bot) SendMessage(body SendMessageRequest) (*Message, error) {
	if body.ParseMode != "" && !body.ParseMode.isValid() {
		return nil, fmt.Errorf("unsupported parse_mode: %s", body.ParseMode)
	}

	res, err := bot.post("sendMessage", body)
	if err != nil {
		return nil, err
	}

	var msg Message
	err = json.Unmarshal(res.Result, &msg)

	return &msg, err
}

// TODO: reply_markup can be InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
type SendPhotoRequest struct {
	BusinessConnectionID    string                   `json:"business_connection_id,omitempty"`
	ChatID                  string                   `json:"chat_id"`
	MessageThreadID         int                      `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   int                      `json:"direct_messages_topic_id,omitempty"`
	Photo                   string                   `json:"photo"`
	Caption                 string                   `json:"caption,omitempty"`
	ParseMode               ParseMode                `json:"parse_mode,omitempty"`
	CaptionEntities         []*MessageEntity         `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                     `json:"show_caption_above_media,omitempty"`
	HasSpoiler              bool                     `json:"has_spoiler,omitempty"`
	DisableNotification     bool                     `json:"disable_notification,omitempty"`
	ProtectContent          bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                     `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         string                   `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             ReplyMarkup              `json:"reply_markup,omitempty"`
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
