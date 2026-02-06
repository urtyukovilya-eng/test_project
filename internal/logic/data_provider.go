package logic

import (
	"context"
	"test_project/internal/dto"
)

type SqlDataProvider interface {
	CreateChat(ctx context.Context, title string) (*dto.Chat, error)
	GetMessagesInChatWithLimit(ctx context.Context, chatId, limit int64) ([]dto.Message, error)
	DeleteMessageByChat(ctx context.Context, chatId int64) error
	CreateMessage(ctx context.Context, chatId int64, text string) (*dto.Message, error)
	DeleteChat(ctx context.Context, id int64) error
	GetChat(ctx context.Context, id int64) (*dto.Chat, error)
}
