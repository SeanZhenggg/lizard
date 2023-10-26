package middleware

import (
	errorToolUtil "github.com/SeanZhenggg/go-utils/errortool"
	"github.com/SeanZhenggg/go-utils/logger"
	"lizard/internal/utils/errs"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	RespData   = "Resp_Data"
	RespStatus = "Resp_Status"
)

type IResponseMiddleware interface {
	Handle(ctx *gin.Context)
}

func ProvideResponseMiddleware(logger logger.ILogger) IResponseMiddleware {
	return &respMiddleware{
		logger,
	}
}

type respMiddleware struct {
	logger logger.ILogger
}

func (respMw *respMiddleware) Handle(ctx *gin.Context) {
	// before request

	ctx.Next()

	// after request
	respMw.standardResponse(ctx)
}

func (respMw *respMiddleware) generateStandardResponse(ctx *gin.Context) response {
	status := ctx.GetInt(RespStatus)
	data := ctx.MustGet(RespData)
	var code int
	var message string

	if status >= http.StatusBadRequest {
		if err, ok := data.(error); ok {
			if parsed, ok := errorToolUtil.ParseError(err); ok {
				code = parsed.GetCode()
				message = parsed.GetMessage()
			} else {
				err, _ := errorToolUtil.ParseError(errs.CommonErr.UnknownError)
				code = err.GetCode()
				message = err.GetMessage()
			}
			data = nil

			respMw.logger.Error(err)
		}
	}

	return response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (respMw *respMiddleware) standardResponse(ctx *gin.Context) {
	response := respMw.generateStandardResponse(ctx)

	respStatus := ctx.GetInt(RespStatus)

	ctx.JSON(
		respStatus,
		response,
	)
}

func SetResp(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.Set(RespStatus, statusCode)
	ctx.Set(RespData, data)
}

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
