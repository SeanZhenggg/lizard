package errs

import (
	"github.com/SeanZhenggg/go-utils/errortool"
)

var (
	Define    = errortool.ProvideDefine()
	CommonErr = Define.Plugin(ProvideCommonError).(*commonError)
	DbErr     = Define.Plugin(ProvideDBError).(*dbError)
)
