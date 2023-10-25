package errs

const (
	ReqGroupCode int = 2
)

func ProvideReqError() *reqError {
	group := Define.GenErrorGroup(ReqGroupCode)

	return &reqError{
		AuthFailedError: group.GenError(1, "權限驗證失敗"),
	}
}

type reqError struct {
	AuthFailedError error
}
