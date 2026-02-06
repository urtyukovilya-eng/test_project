package logic

import (
	"context"
	"test_project/internal/dto"
	"test_project/internal/logic"

	"go.uber.org/zap"
)

type Logic struct {
	l            *zap.Logger
	dataProvider logic.SqlDataProvider
}

func NewLogic(l *zap.Logger,
	dataProvider logic.SqlDataProvider,
) *Logic {
	return &Logic{
		l:            l,
		dataProvider: dataProvider,
	}
}

func (l *Logic) NewChat(ctx context.Context, title string) (*dto.Chat, error) {
	chat, err := l.dataProvider.CreateChat(ctx, title)
	if err != nil {
		l.l.Sugar().Errorln(err)
		return nil, err
	}
	l.l.Sugar().Debugln(chat)
	return chat, nil
}

func (l *Logic) NewMessage(ctx context.Context, chatID int64, text string) (*dto.Message, error) {
	chat, err := l.dataProvider.GetChat(ctx, chatID)
	if err != nil {
		l.l.Sugar().Errorln(err)
		return nil, err
	}
	if chat == nil {
		return nil, dto.ErrChatNotExist
	}
	message, err := l.dataProvider.CreateMessage(ctx, chatID, text)
	if err != nil {
		l.l.Sugar().Errorln(err)
		return nil, err
	}
	l.l.Sugar().Debugln(message)
	return message, nil
}

func (l *Logic) DeleteChat(ctx context.Context, chatID int64) error {
	err := l.dataProvider.DeleteMessageByChat(ctx, chatID)
	if err != nil {
		l.l.Sugar().Errorln(err)
		return err
	}
	err = l.dataProvider.DeleteChat(ctx, chatID)
	if err != nil {
		l.l.Sugar().Errorln(err)
		return err
	}
	return nil
}

func (l *Logic) GetChat(ctx context.Context, chatID, limit int64) (*dto.Chat, []dto.Message, error) {
	messages, err := l.dataProvider.GetMessagesInChatWithLimit(ctx, chatID, limit)
	if err != nil {
		l.l.Sugar().Errorln(err)
		return nil, nil, err
	}
	chat, err := l.dataProvider.GetChat(ctx, chatID)
	if err != nil {
		l.l.Sugar().Errorln(err)
		return nil, nil, err
	}
	if chat == nil {
		return nil, nil, nil
	}
	return chat, messages, err
}
