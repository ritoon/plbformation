package db

import "errors"

type ErrorDb struct {
	Message      string
	Params       map[string]interface{}
	FunctionName string
	IsCritial    bool
}

func (e *ErrorDb) Error() string {
	return e.Message
}

func NewErrorDb(message string, params map[string]interface{}, functionName string, isCritical bool) *ErrorDb {
	return &ErrorDb{
		Message:      message,
		Params:       params,
		FunctionName: functionName,
		IsCritial:    isCritical,
	}
}

var (
	ErrProductNotFound = NewErrorDb("product not found", nil, "GetByID Product", false)
	ErrUserNotFound    = NewErrorDb("user not found", nil, "GetByID User", false)
	ErrorDbConnection  = errors.New("db: db connexion")
)

func IsErrorDb(err error) bool {
	switch err.(type) {
	case *ErrorDb:
		return true
	default:
		return false
	}
}
