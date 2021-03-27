package common_test

import (
	"testing"

	"github.com/dundee/go-qrcode-payment/common"
	"github.com/stretchr/testify/assert"
)

func TestTrimToLength(t *testing.T) {
	s := common.TrimToLength("12345", 3)
	assert.Equal(t, "123", s)

	s = common.TrimToLength("321", 5)
	assert.Equal(t, "321", s)
}
