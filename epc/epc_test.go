package epc_test

import (
	"testing"

	"github.com/dundee/qrpay/epc"
	"github.com/stretchr/testify/assert"
)

func TestSettingLongIBAN(t *testing.T) {
	p := epc.NewEpcPayment()
	p.SetRecipientName("XX")
	p.IBAN = "12345678901234567890123456789012345678901234567890" // 50 chars

	s, _ := p.GenerateString()
	// trimmed to 46 chars
	assert.Equal(t, "BCD\n002\n1\nSCT\n\nXX\n1234567890123456789012345678901234", s)
}

func TestGenerateString(t *testing.T) {
	p := epc.NewEpcPayment()
	p.SetRecipientName("Red Cross")
	p.SetIBAN("CZ5855000000001265098001")
	p.SetBIC("RZBCCZPP")
	p.SetAmount("100.0")

	s, _ := p.GenerateString()
	assert.Equal(t, "BCD\n002\n1\nSCT\nRZBCCZPP\nRed Cross\nCZ5855000000001265098001\n100.0", s)
}

func TestDEPayment(t *testing.T) {
	p := epc.NewEpcPayment()
	p.SetRecipientName("Franz Mustermänn")
	p.SetIBAN("DE71110220330123456789")
	p.SetBIC("BHBLDEHHXXX")
	p.SetCurrency("EUR")
	p.SetAmount("10.8")

	s, _ := p.GenerateString()
	assert.Equal(t, "BCD\n002\n1\nSCT\nBHBLDEHHXXX\nFranz Mustermänn\nDE71110220330123456789\n10.8", s)
}

func TestOtherParams(t *testing.T) {
	p := epc.NewEpcPayment()
	p.SetIBAN("CZ5855000000001265098001")
	p.SetSenderReference("111111")
	p.SetMessage("M")
	p.SetRecipientName("go")
	p.SetPurpose("GDDS")

	s, _ := p.GenerateString()
	assert.Equal(
		t,
		"BCD\n002\n1\nSCT\n\ngo\nCZ5855000000001265098001\n\nGDDS\n111111\n\nM",
		s,
	)
}

func TestGenerateStringWithoutIBAN(t *testing.T) {
	p := epc.NewEpcPayment()
	p.SetRecipientName("xxx")
	s, err := p.GenerateString()
	assert.Equal(t, "", s)
	assert.Equal(t, "IBAN is mandatory", err.Error())
}

func TestGenerateStringWithoutRecipient(t *testing.T) {
	p := epc.NewEpcPayment()
	s, err := p.GenerateString()
	assert.Equal(t, "", s)
	assert.Equal(t, "name of the beneficiary is mandatory", err.Error())
}
