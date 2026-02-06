package logic

import (
	"context"
	"reflect"
	"test_project/internal/dto"
	"testing"
)

func TestLogic_NewChat(t *testing.T) {
	type args struct {
		ctx   context.Context
		title string
	}
	tests := []struct {
		name    string
		l       *Logic
		args    args
		want    *dto.Chat
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.NewChat(tt.args.ctx, tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("Logic.NewChat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Logic.NewChat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogic_NewMessage(t *testing.T) {
	type args struct {
		ctx    context.Context
		chatID int64
		text   string
	}
	tests := []struct {
		name    string
		l       *Logic
		args    args
		want    *dto.Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.NewMessage(tt.args.ctx, tt.args.chatID, tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Logic.NewMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Logic.NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogic_DeleteChat(t *testing.T) {
	type args struct {
		ctx    context.Context
		chatID int64
	}
	tests := []struct {
		name    string
		l       *Logic
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.l.DeleteChat(tt.args.ctx, tt.args.chatID); (err != nil) != tt.wantErr {
				t.Errorf("Logic.DeleteChat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_GetChat(t *testing.T) {
	type args struct {
		ctx    context.Context
		chatID int64
		limit  int64
	}
	tests := []struct {
		name    string
		l       *Logic
		args    args
		want    *dto.Chat
		want1   []dto.Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.l.GetChat(tt.args.ctx, tt.args.chatID, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("Logic.GetChat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Logic.GetChat() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Logic.GetChat() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
