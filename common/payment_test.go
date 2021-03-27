package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrongIban(t *testing.T) {
	defer func() {
		err := recover()
		assert.NotNil(t, err)
		assert.Equal(t, "iban: iban too short", err.(error).Error())
	}()

	p := NewPayment()
	p.SetIBAN("111")
}
