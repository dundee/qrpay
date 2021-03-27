package common

import (
	"github.com/jbub/banking/iban"
	"github.com/jbub/banking/swift"
)

type Payment struct {
	IBAN      string
	BIC       string
	Amount    string
	Currency  string
	Reference string
	Recipient string
	Msg       string
}

func NewPayment() *Payment {
	return &Payment{}
}

func (p *Payment) SetIBAN(value string) {
	iban.MustParse(value)
	p.IBAN = value
}

func (p *Payment) SetBIC(value string) {
	swift.MustParse(value)
	p.BIC = value
}

func (p *Payment) SetAmount(value string) {
	p.Amount = value
}

func (p *Payment) SetCurrency(value string) {
	p.Currency = value
}

func (p *Payment) SetSenderReference(value string) {
	p.Reference = value
}

func (p *Payment) SetRecipientName(value string) {
	p.Recipient = value
}

func (p *Payment) SetMessage(value string) {
	p.Msg = value
}
