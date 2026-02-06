package dto

import (
	"errors"
	"time"
)

const (
	DefaultLimit     = 20
	MaxLimit         = 100
	MaxTitleLength   = 200
	MaxMessageLength = 5000
)

type Chat struct {
	Id        uint64    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type Message struct {
	Id        uint64    `json:"id"`
	ChatId    uint64    `json:"chat_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

var ErrChatNotExist = errors.New("chat not exist")
