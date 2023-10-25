package errs

const (
	DBGroupCode int = 9
)

func ProvideDBError() *dbError {
	group := Define.GenErrorGroup(DBGroupCode)

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
