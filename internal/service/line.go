package service

import (
	"context"
	"github.com/SeanZhenggg/go-utils/logger"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"golang.org/x/xerrors"
	"lizard/internal/constant"
	"lizard/internal/model/bo"
	"log"
)

type IMessageSrv interface {
	PushMessage(ctx context.Context, cond *bo.SendMessage) error
}

type lineSrv struct {
	bot    *linebot.Client
	logger logger.ILogger
}

func ProvideLineSrv(logger logger.ILogger) IMessageSrv {
	bot, err := linebot.New(constant.ChannelSecret, constant.ChannelAccessToken)
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
