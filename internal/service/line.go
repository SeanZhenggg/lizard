package service

import (
	"context"
	"github.com/SeanZhenggg/go-utils/logger"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"golang.org/x/xerrors"
	"lizard/internal/constant"
	"lizard/internal/model/bo"
	"log"
	"math"
)

const (
	MAX_SEND_LENGTH = 5
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
	totalSendLen := len(cond.Messages)
	interval := int(math.Ceil(float64(totalSendLen) / MAX_SEND_LENGTH))

	for i := 0; i < interval; i++ {
		var messages []linebot.SendingMessage
		if len(cond.Messages) < (i+1)*MAX_SEND_LENGTH {
			messages = cond.Messages[i*MAX_SEND_LENGTH:]
		} else {
			messages = cond.Messages[i*MAX_SEND_LENGTH : (i+1)*MAX_SEND_LENGTH]
		}

		if _, err := srv.bot.PushMessage(cond.To, messages...).Do(); err != nil {
			return xerrors.Errorf("lineSrv PushMessage error : %w\n", err)
		}
	}

	return nil
}
