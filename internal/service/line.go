package service

import (
	"context"
	"github.com/SeanZhenggg/go-utils/logger"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"golang.org/x/xerrors"
	"lizard/internal/model/bo"
	"log"
)

const (
	CHANNEL_SECRET       = "d48ec0332fb59de64035f2c555bb9995"
	CHANNEL_ACCESS_TOKEN = "aAV/cg5zZd3nY19Po6aawPs7C1vBotu+AJHs0wMdndKbrm5XPqQ4XOYUa1QSYrDBfvLK+t0JU8v1cf1BQbZvepXinztFKns2xa79JZEbTbRbRPUqcuZXbYY5FF5eAEyJmAR8msweC0yumXr11Pz62QdB04t89/1O/w1cDnyilFU="
)

type IMessageSrv interface {
	PushMessage(ctx context.Context, cond *bo.SendMessage) error
}

type lineSrv struct {
	bot    *linebot.Client
	logger logger.ILogger
}

func ProvideLineSrv(logger logger.ILogger) IMessageSrv {
	bot, err := linebot.New(CHANNEL_SECRET, CHANNEL_ACCESS_TOKEN)
	if err != nil {
		log.Fatal(xerrors.Errorf("CreateNewBot error : %w", err))
	}

	return &lineSrv{
		bot:    bot,
		logger: logger,
	}

}

func (srv *lineSrv) PushMessage(ctx context.Context, cond *bo.SendMessage) error {
	_, err := srv.bot.PushMessage(cond.To, cond.Messages...).Do()
	if err != nil {
		return xerrors.Errorf("lineSrv PushMessage error : %w\n", err)
	}

	return nil
}
