package errs

import (
	"github.com/SeanZhenggg/go-utils/errortool"
)

var ()

var (
	Define    = errortool.ProvideDefine()
	CommonErr = ProvideCommonError()
	DbErr     = ProvideDBError()
)
