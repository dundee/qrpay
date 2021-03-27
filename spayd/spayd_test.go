package spayd_test

import (
	"fmt"
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

func ExampleSpaydPayment_GenerateString() {
	p := spayd.NewSpaydPayment()
	p.SetIBAN("CZ5855000000001265098001")
	p.SetBIC("RZBCCZPP")
	p.SetAmount("100.0")

	s, _ := p.GenerateString()
	fmt.Println(s)
	// Output: SPD*1.0*ACC:CZ5855000000001265098001+RZBCCZPP*AM:100.0*
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
	p.SetPaymentType("P2P")
	p.SetSenderReference("111")
	p.SetMessage("M")
	p.SetRecipientName("go")
	p.SetNofificationType('E')
	p.SetNotificationValue("daniel@milde.cz")

	s, _ := p.GenerateString()
	assert.Equal(
		t,
		"SPD*1.0*ACC:CZ5855000000001265098001*RF:111*RN:GO*DT:20211224*MSG:M*PT:P2P*NT:E*NT:DANIEL@MILDE.CZ*",
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

func TestWrongNotificationType(t *testing.T) {
	defer func() {
		err := recover()
		assert.NotNil(t, err)
		assert.Equal(t, "nofification type 'X' is not supported (E | P)", err)
	}()

	p := spayd.NewSpaydPayment()
	p.SetNofificationType('X')
}

func TestAsteriskInValue(t *testing.T) {
	p := spayd.NewSpaydPayment()
	p.SetIBAN("CZ5855000000001265098001")
	p.SetMessage("aaa*bbb")

	s, err := p.GenerateString()
	assert.Nil(t, err)
	assert.Equal(t, "SPD*1.0*ACC:CZ5855000000001265098001*MSG:AAA%2ABBB*", s)
}
