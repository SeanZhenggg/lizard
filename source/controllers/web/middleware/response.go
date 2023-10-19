package middleware

import (
	"errors"
	errorToolUtil "github.com/SeanZhenggg/go-utils/errortool"
	"lizard/source/utils/errortool"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	RespData   = "Resp_Data"
	RespStatus = "Resp_Status"
)

type IResponseMiddleware interface {
	ResponseHandler(ctx *gin.Context)
}

func ProvideResponseMiddleware() IResponseMiddleware {
	return &ResponseMiddleware{}
}

type ResponseMiddleware struct{}

func (respMw *ResponseMiddleware) ResponseHandler(ctx *gin.Context) {
	// before request

	ctx.Next()

	// after request
	respMw.standardResponse(ctx)
}

func (respMw *ResponseMiddleware) generateStandardResponse(ctx *gin.Context) response {
	status := ctx.GetInt(RespStatus)
	data := ctx.MustGet(RespData)
	var code int
	var message string

	if status >= http.StatusBadRequest {
		var err error
		if errors.As(data.(error), &err) {
			if parsed, ok := errorToolUtil.ParseError(err); ok {
				code = parsed.GetCode()
				message = parsed.GetMessage()
			} else {
				err, _ := errorToolUtil.ParseError(errortool.CommonErr.UnknownError)
				code = err.GetCode()
				message = err.GetMessage()
			}
			data = nil
		}
	}

	return response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (respMw *ResponseMiddleware) standardResponse(ctx *gin.Context) {
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
