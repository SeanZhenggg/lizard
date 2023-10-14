package middleware

import (
	"fmt"

	logUtils "github.com/SeanZhenggg/go-utils/logger"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	HEADER_AUTHORIZATION = "Authorization"
	ATHORIZATION_PREFIX  = "Bearer "
)

type IAuthMiddleware interface {
	AuthValidationHandler(ctx *gin.Context)
}

func ProvideAuthMiddleware(logger logUtils.ILogger) IAuthMiddleware {
	return &AuthMiddleware{
		logger: logger,
	}
}

type AuthMiddleware struct {
	logger logUtils.ILogger
}

func (authMw *AuthMiddleware) AuthValidationHandler(ctx *gin.Context) {
	// before request
	tokenStr := ctx.GetHeader(HEADER_AUTHORIZATION)

	token, _ := strings.CutPrefix(tokenStr, ATHORIZATION_PREFIX)

	authMw.logger.Info(fmt.Sprintf("token %v\n", token))

	ctx.Next()

	// after request
}
