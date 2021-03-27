package spayd_test

import (
	"testing"
	"time"

	"github.com/dundee/go-qrcode-payment/spayd"
	"github.com/stretchr/testify/assert"
)

func TestSettingLongAcc(t *testing.T) {
	p := spayd.NewSpaydPayment()
	p.IBAN = "12345678901234567890123456789012345678901234567890" // 50 chars

	s, _ := p.GenerateString()
	// trimmed to 46 chars
	assert.Equal(t, "SPD*1.0*ACC:1234567890123456789012345678901234567890123456*", s)
}

func TestCZPayment(t *testing.T) {
	p := spayd.NewSpaydPayment()
	p.SetIBAN("CZ5855000000001265098001")
	p.SetBIC("RZBCCZPP")
	p.SetAmount("100.0")

	s, _ := p.GenerateString()
	assert.Equal(t, "SPD*1.0*ACC:CZ5855000000001265098001+RZBCCZPP*AM:100.0*", s)
}

func TestDEPayment(t *testing.T) {
	p := spayd.NewSpaydPayment()
	p.SetIBAN("DE71110220330123456789")
	p.SetBIC("BHBLDEHHXXX")
	p.SetCurrency("EUR")
	p.SetAmount("10.8")

	s, _ := p.GenerateString()
	assert.Equal(t, "SPD*1.0*ACC:DE71110220330123456789+BHBLDEHHXXX*AM:10.8*CC:EUR*", s)
}

func TestOtherParams(t *testing.T) {
	p := spayd.NewSpaydPayment()
	p.SetIBAN("CZ5855000000001265098001")
	p.SetDate(time.Date(2021, 12, 24, 0, 0, 0, 0, time.UTC))
	p.SetMessage("M")
	p.SetRecipientName("go")
	p.SetNofificationType('E')
	p.SetNotificationValue("daniel@milde.cz")

	s, _ := p.GenerateString()
	assert.Equal(
		t,
		"SPD*1.0*ACC:CZ5855000000001265098001*RN:GO*DT:20211224*MSG:M*NT:E*NT:DANIEL@MILDE.CZ*",
		s,
	)
}

func TestExtended(t *testing.T) {
	p := spayd.NewSpaydPayment()
	p.SetIBAN("CZ5855000000001265098001")
	p.SetExtendedAttribute("vs", "1234567890")

	s, _ := p.GenerateString()
	assert.Equal(t, "SPD*1.0*ACC:CZ5855000000001265098001*X-VS:1234567890*", s)
}

func TestGenerateStringWithourIBAN(t *testing.T) {
	p := spayd.NewSpaydPayment()
	s, err := p.GenerateString()
	assert.Equal(t, "", s)
	assert.Equal(t, "IBAN is mandatory", err.Error())
}
