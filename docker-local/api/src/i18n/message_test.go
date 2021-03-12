package i18n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJaMessage(t *testing.T) {
	actual := Message("error.unavailable")
	assert.Equal(t, actual, "現在サービスは利用できません。しばらく時間を置いてから、もう一度お試しください。")
}
