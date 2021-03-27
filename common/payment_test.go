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

func TestMethods(t *testing.T) {
	p := NewPayment()
	p.SetIBAN("CZ5855000000001265098001")
	p.SetBIC("BHBLDEHHXXX")
	p.SetCurrency("EUR")
	p.SetAmount("10.8")
	p.SetMessage("M")
	p.SetRecipientName("go")
	p.SetSenderReference("RF:111")

	assert.Equal(t, "CZ5855000000001265098001", p.IBAN)
	assert.Equal(t, "BHBLDEHHXXX", p.BIC)
	assert.Equal(t, "EUR", p.Currency)
	assert.Equal(t, "10.8", p.Amount)
	assert.Equal(t, "M", p.Msg)
	assert.Equal(t, "go", p.Recipient)
	assert.Equal(t, "RF:111", p.Reference)
}
