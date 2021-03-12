// Package misc defines miscellaneous functions
package misc

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/supinf/elasticsearch-example/api/src/i18n"
)

func TestServiceUnavailable(t *testing.T) {
	expect := &Error{
		StatusCode:         http.StatusServiceUnavailable,
		ErrorCodeForClient: "E100000",
		ErrorMsgForClient:  i18n.Message("error.unavailable"),
	}
	actual := ServiceUnavailable
	assert.Equal(t, actual, expect)
}
