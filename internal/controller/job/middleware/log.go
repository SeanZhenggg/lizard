package middleware

import (
	"github.com/SeanZhenggg/go-utils/logger"
	"lizard/internal/utils/cronjob"
)

const (
	RespData   = "Resp_Data"
	RespStatus = "Resp_Status"
)

type IJobLogMiddleware interface {
	Handle(ctx *cronjob.Context)
}

func ProvideJobLogMiddleware(logger logger.ILogger) IJobLogMiddleware {
	return &logMiddleware{
		logger,
	}
}

type logMiddleware struct {
	logger logger.ILogger
}

func (respMw *logMiddleware) Handle(ctx *cronjob.Context) {
	// before request

	ctx.Next()

	// after request
	respMw.log(ctx)
}

func (respMw *logMiddleware) log(ctx *cronjob.Context) {
	data, existed := ctx.Get(RespData)

	if !existed {
		return
	}

	if err, ok := data.(error); ok {
		respMw.logger.Error(err)
	}
}
