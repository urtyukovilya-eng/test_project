package logic

import (
	"context"
	"test_project/internal/dto"
)

type Logic interface {
	NewChat(ctx context.Context, title string) (*dto.Chat, error)
	NewMessage(ctx context.Context, chatID int64, text string) (*dto.Message, error)
	GetChat(ctx context.Context, chatID, limit int64) (*dto.Chat, []dto.Message, error)
	DeleteChat(ctx context.Context, chatID int64) error
}
