// Package misc defines miscellaneous functions
package misc

import (
	"github.com/supinf/elasticsearch-example/api/src/i18n"
)

// ---------------------------------------------------------------------
//  ユーザ、またはクライアントアプリに返すエラー
// ---------------------------------------------------------------------

// Error エラー発生に関する情報
type Error struct {
	StatusCode         int
	ErrorCodeForClient string
	ErrorMsgForClient  string
}

func newError(httpStatusCode int, errCode, errMsg string) *Error {
	return &Error{
		StatusCode:         httpStatusCode,
		ErrorCodeForClient: errCode,
		ErrorMsgForClient:  i18n.Message(errMsg),
	}
}

// Equals 同じエラーコードのエラーかを判定
func (e *Error) Equals(err *Error) bool {
	if err == nil {
		return false
	}
	if e == nil {
		return false
	}
	return e.ErrorCodeForClient == err.ErrorCodeForClient
}

// IsStatus
func (e Error) IsStatus(code int) bool {
	return e.StatusCode == code
}
