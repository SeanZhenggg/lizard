package middleware

import (
	"fmt"
	"github.com/SeanZhenggg/go-utils/logger"
	"lizard/internal/utils/cronjob"
	"lizard/internal/utils/log"
	"runtime"
)

const (
	LogFuncName = "FUNC_NAME"
	LogError    = "ERROR"
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

func (mw *logMiddleware) Handle(ctx *cronjob.Context) {
	// before request
	defer log.CallTotalDurationLog(func(spentTime string) {
		funcName, _ := ctx.Get(LogFuncName)
		mw.logger.Info(fmt.Sprintf("job func \"%s\" execution completed, total duration : %s", funcName, spentTime))
	})()

	ctx.Next()

	// after request
	mw.log(ctx)
}

func (mw *logMiddleware) log(ctx *cronjob.Context) {
	data, existed := ctx.Get(LogError)

	if !existed {
		return
	}

	if err, ok := data.(error); ok {
		mw.logger.Error(err)
	}
}

func SetJobError(ctx *cronjob.Context, data error) {
	ctx.Set(LogError, data)
}

func SetJobFunc(ctx *cronjob.Context) {
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()

	ctx.Set(LogFuncName, funcName)
}
