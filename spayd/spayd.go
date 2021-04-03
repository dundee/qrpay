package spayd

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/dundee/qrpay/base"
)

const SpaydHeader = "SPD*1.0*"

type SpaydPayment struct {
	*base.Payment
	Date        time.Time
	PaymentType string
	NotifType   rune
	NotifValue  string
	Extended    map[string]string
}

func NewSpaydPayment() *SpaydPayment {
	return &SpaydPayment{
		Payment: &base.Payment{
			Errors: make(map[string]error),
		},
		Extended: make(map[string]string),
	}
}

func (s *SpaydPayment) SetDate(value time.Time) {
	s.Date = value
}

func (s *SpaydPayment) SetPaymentType(value string) {
	s.PaymentType = value
}

func (s *SpaydPayment) SetNofificationType(value rune) {
	if value != 'P' && value != 'E' {
		panic("nofification type '" + string(value) + "' is not supported (E | P)")
	}
	s.NotifType = value
}

func (s *SpaydPayment) SetExtendedAttribute(name string, value string) {
	s.Extended[name] = value
}

func (s *SpaydPayment) SetNotificationValue(value string) {
	s.NotifValue = value
}

func (s *SpaydPayment) GenerateString() (string, error) {
	res := strings.Builder{}
	res.WriteString(SpaydHeader)

	if s.IBAN == "" {
		return "", errors.New("IBAN is mandatory")
	}

	acc := s.IBAN
	if s.BIC != "" {
		acc += "+" + s.BIC
	}
	res.WriteString("ACC:" + base.TrimToLength(convertValue(acc), 46) + "*")

	if s.Amount != "" {
		res.WriteString("AM:" + base.TrimToLength(convertValue(s.Amount), 10) + "*")
	}

	if s.Currency != "" {
		res.WriteString("CC:" + base.TrimToLength(convertValue(s.Currency), 3) + "*")
	}

	if s.Reference != "" {
		res.WriteString("RF:" + base.TrimToLength(convertValue(s.Reference), 16) + "*")
	}

	if s.Recipient != "" {
		res.WriteString("RN:" + base.TrimToLength(convertValue(s.Recipient), 35) + "*")
	}

	if s.Date.Year() > 1 {
		year, month, day := s.Date.Date()
		res.WriteString(fmt.Sprintf("DT:%4d%d%d*", year, month, day))
	}

	if s.Msg != "" {
		res.WriteString("MSG:" + base.TrimToLength(convertValue(s.Msg), 60) + "*")
	}

	if s.PaymentType != "" {
		res.WriteString("PT:" + base.TrimToLength(convertValue(s.PaymentType), 3) + "*")
	}

	if s.NotifType > 0 {
		res.WriteString("NT:" + string(s.NotifType) + "*")
	}
	if s.NotifValue != "" {
		res.WriteString("NT:" + base.TrimToLength(convertValue(s.NotifValue), 320) + "*")
	}

	for name := range s.Extended {
		res.WriteString("X-" + convertValue(name) + ":" + convertValue(s.Extended[name]) + "*")
	}

	return res.String(), nil
}

func convertValue(value string) string {
	value = strings.ToUpper(value)
	value = url.PathEscape(value)
	return value
}
