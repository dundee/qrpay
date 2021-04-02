package epc

import (
	"errors"
	"strings"

	"github.com/dundee/go-qrcode-payment/common"
)

const EpcHeader = `BCD
002
1
SCT
`

type EpcPayment struct {
	*common.BasePayment
	Purpose string
}

func NewEpcPayment() *EpcPayment {
	return &EpcPayment{
		BasePayment: &common.BasePayment{
			Errors: make(map[string]error),
		},
	}
}

func (p *EpcPayment) SetPurpose(value string) {
	p.Purpose = value
}

func (p *EpcPayment) GenerateString() (string, error) {
	res := strings.Builder{}
	res.WriteString(EpcHeader)

	if p.BIC != "" {
		res.WriteString(common.TrimToLength(p.BIC, 11))
	}
	res.WriteString("\n")

	if p.Recipient == "" {
		return "", errors.New("name of the beneficiary is mandatory")
	}
	res.WriteString(common.TrimToLength(p.Recipient, 70) + "\n")

	if p.IBAN == "" {
		return "", errors.New("IBAN is mandatory")
	}
	res.WriteString(common.TrimToLength(p.IBAN, 34) + "\n")

	if p.Amount != "" {
		res.WriteString(common.TrimToLength(p.Amount, 12))
	}
	res.WriteString("\n")

	if p.Purpose != "" {
		res.WriteString(common.TrimToLength(p.Purpose, 4))
	}
	res.WriteString("\n")

	if p.Reference != "" {
		res.WriteString(common.TrimToLength(p.Reference, 4))
	}
	res.WriteString("\n\n")

	if p.Msg != "" {
		res.WriteString(common.TrimToLength(p.Msg, 70))
	}

	return strings.TrimSpace(res.String()), nil
}

var _ common.Payment = (*EpcPayment)(nil)
