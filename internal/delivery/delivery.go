package delivery

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"test_project/internal/dto"
	"test_project/internal/logic"

	"go.uber.org/zap"
)

type Delivery struct {
	l            *zap.Logger
	internaLogic logic.Logic
}

func New(l *zap.Logger, internaLogic logic.Logic) *Delivery {
	return &Delivery{
		l:            l,
		internaLogic: internaLogic,
	}
}

func (d *Delivery) newChat(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	body := make(map[string]string)
	request, err := io.ReadAll(r.Body)
	if err != nil {
		d.l.Sugar().Errorln(err)
	}
	err = json.Unmarshal(request, &body)
	if err != nil {
		d.l.Sugar().Errorln(err)
	}
	title := body["title"]
	title = strings.TrimSpace(title)
	if title == "" || len([]rune(title)) > dto.MaxTitleLength {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	chat, err := d.internaLogic.NewChat(ctx, title)
	if err != nil {
		d.l.Sugar().Errorln(err)
	}

	response, err := json.Marshal(chat)
	if err != nil {
		d.l.Sugar().Errorln(err)
	}
	_, err = w.Write(response)
	if err != nil {
		d.l.Sugar().Errorln(err)
	}
}

func (d *Delivery) newMewssage(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	body := make(map[string]string)
	request, err := io.ReadAll(r.Body)
	if err != nil {
		d.l.Sugar().Errorln(err)
	}
	err = json.Unmarshal(request, &body)
	if err != nil {
		d.l.Sugar().Errorln(err)
	}
	chatIdString := r.PathValue("id")
	chatId, err := strconv.Atoi(chatIdString)
	if err != nil {
		d.l.Sugar().Errorln(err)
	}
	d.l.Sugar().Debugln(chatId, body)
	txt := body["text"]
	if txt == "" || len([]rune(txt)) > dto.MaxMessageLength {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	mes, err := d.internaLogic.NewMessage(ctx, int64(chatId), txt)
	if err != nil {
		if errors.Is(err, dto.ErrChatNotExist) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		d.l.Sugar().Errorln(err)
	}

	response, err := json.Marshal(mes)
	if err != nil {
		d.l.Sugar().Errorln(err)
	}
	_, err = w.Write(response)
	if err != nil {
		d.l.Sugar().Errorln(err)
	}
}

func (d *Delivery) deleteChat(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	chatIdString := r.PathValue("id")
	chatId, err := strconv.Atoi(chatIdString)
	if err != nil {
		d.l.Sugar().Errorln(err)
	}
	d.l.Sugar().Debugln(chatId)

	err = d.internaLogic.DeleteChat(ctx, int64(chatId))
	if err != nil {
		d.l.Sugar().Errorln(err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (d *Delivery) getChat(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	chatIdString := r.PathValue("id")
	chatId, err := strconv.Atoi(chatIdString)
	if err != nil {
		d.l.Sugar().Errorln(err)
	}
	d.l.Sugar().Debugln(chatId)
	limStr := r.URL.Query().Get("limit")
	var lim int
	if limStr != "" {
		lim, err = strconv.Atoi(limStr)
		if err != nil {
			d.l.Sugar().Errorln(err)
		}
	} else {
		lim = dto.DefaultLimit
	}
	if lim > dto.MaxLimit {
		lim = dto.MaxLimit
	}
	d.l.Sugar().Debugln(lim)

	chat, messages, err := d.internaLogic.GetChat(ctx, int64(chatId), int64(lim))
	if err != nil {
		d.l.Sugar().Errorln(err)
	}
	if chat == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	type responseStruct struct {
		Chat     dto.Chat      `json:"chat"`
		Messages []dto.Message `json:"messages"`
	}
	response, err := json.Marshal(responseStruct{Chat: *chat, Messages: messages})
	if err != nil {
		d.l.Sugar().Errorln(err)
	}
	_, err = w.Write(response)
	if err != nil {
		d.l.Sugar().Errorln(err)
	}
}
