package errortool

import (
	"github.com/SeanZhenggg/go-utils/errortool"
)

const (
	DBGroupCode int = 9
)

func ProvideDBError(groups errortool.IGroupRepo, codes errortool.ICodeRepo) interface{} {
	group := Define.GenGroup(DBGroupCode)

	return &dbError{
		TestError: group.GenError(1, "Test error"),
	}
}

var (
// TODO
)

type dbError struct {
	// TODO
	TestError error
}

func ParseDBError(err error) error {
	// TODO
	return err
}
