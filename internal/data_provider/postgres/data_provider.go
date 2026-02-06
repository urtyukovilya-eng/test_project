package repository

import (
	"context"
	"test_project/internal/dto"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SqlRepository struct {
	db *gorm.PreparedStmtDB
	l  *zap.Logger
}

func New(l *zap.Logger, db *gorm.PreparedStmtDB) *SqlRepository {
	return &SqlRepository{
		l:  l,
		db: db,
	}
}

func (r *SqlRepository) CreateChat(ctx context.Context, title string) (*dto.Chat, error) {
	r.l.Sugar().Infoln("start create chat")
	row, err := r.db.QueryContext(ctx, "INSERT INTO public.chat (title,created_at) VALUES ($1,$2) RETURNING *;", title, time.Now().UTC())
	if err != nil {
		r.l.Sugar().Errorln(err)
		return nil, err
	}
	var chat dto.Chat
	row.Next()
	err = row.Scan(&chat.Id, &chat.Title, &chat.CreatedAt)
	if err != nil {
		r.l.Sugar().Errorln(err)
		return nil, err
	}
	return &chat, nil
}

func (r *SqlRepository) DeleteChat(ctx context.Context, id int64) error {
	r.l.Sugar().Infoln("start delete chat")
	_, err := r.db.ExecContext(ctx, "DELETE FROM public.chat WHERE id=$1;", id)
	if err != nil {
		r.l.Sugar().Errorln(err)
		return err
	}
	return nil
}

func (r *SqlRepository) CreateMessage(ctx context.Context, chatId int64, text string) (*dto.Message, error) {
	r.l.Sugar().Infoln("start create message")
	row, err := r.db.QueryContext(ctx,
		"INSERT INTO public.message (chat_id,text,created_at) VALUES ($1,$2,$3) RETURNING *;",
		chatId, text, time.Now().UTC())
	if err != nil {
		r.l.Sugar().Errorln(err)
		return nil, err
	}
	var mes dto.Message
	row.Next()
	err = row.Scan(&mes.Id, &mes.ChatId, &mes.Text, &mes.CreatedAt)
	if err != nil {
		r.l.Sugar().Errorln(err)
		return nil, err
	}
	return &mes, nil
}

func (r *SqlRepository) DeleteMessageByChat(ctx context.Context, chatId int64) error {
	r.l.Sugar().Infoln("start delete messages")
	_, err := r.db.ExecContext(ctx, "DELETE FROM public.message WHERE chat_id=$1;", chatId)
	if err != nil {
		r.l.Sugar().Errorln(err)
		return err
	}
	return nil
}

func (r *SqlRepository) GetMessagesInChatWithLimit(ctx context.Context, chatId, limit int64) ([]dto.Message, error) {
	row, err := r.db.QueryContext(ctx, "SELECT * FROM public.message WHERE chat_id=$1 ORDER BY created_at DESC LIMIT $2;", chatId, limit)
	if err != nil {
		r.l.Sugar().Errorln(err)
		return nil, err
	}
	mes := make([]dto.Message, 0)
	for row.Next() {
		var message dto.Message
		err = row.Scan(&message.Id, &message.ChatId, &message.Text, &message.CreatedAt)
		if err != nil {
			r.l.Sugar().Errorln(err)
			return nil, err
		}
		mes = append(mes, message)
	}

	return mes, nil
}

func (r *SqlRepository) GetChat(ctx context.Context, chatID int64) (*dto.Chat, error) {
	row, err := r.db.QueryContext(ctx, "SELECT * FROM public.chat WHERE id=$1;", chatID)
	if err != nil {
		r.l.Sugar().Errorln(err)
		return nil, err
	}
	var chat dto.Chat
	hasChat := row.Next()
	if !hasChat {
		return nil, nil
	}
	err = row.Scan(&chat.Id, &chat.Title, &chat.CreatedAt)
	if err != nil {
		r.l.Sugar().Errorln(err)
		return nil, err
	}
	return &chat, nil
}
