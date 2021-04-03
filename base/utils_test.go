package base_test

import (
	"testing"

	"github.com/dundee/qrpay/base"
	"github.com/stretchr/testify/assert"
)

func TestTrimToLength(t *testing.T) {
	s := base.TrimToLength("12345", 3)
	assert.Equal(t, "123", s)

	s = base.TrimToLength("321", 5)
	assert.Equal(t, "321", s)
}
