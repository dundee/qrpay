package common

import (
	"github.com/jbub/banking/iban"
	"github.com/jbub/banking/swift"
)

type BasePayment struct {
	IBAN      string
	BIC       string
	Amount    string
	Currency  string
	Reference string
	Recipient string
	Msg       string
	Errors    map[string]error
}

func NewBasePayment() *BasePayment {
	return &BasePayment{
		Errors: make(map[string]error),
	}
}

func (p *BasePayment) SetIBAN(value string) error {
	_, err := iban.Parse(value)
	if err != nil {
		p.Errors["iban"] = err
	} else {
		p.IBAN = value
	}
	return err
}

func (p *BasePayment) SetBIC(value string) error {
	_, err := swift.Parse(value)
	if err != nil {
		p.Errors["bic"] = err
	} else {
		p.BIC = value
	}
	return err
}

func (p *BasePayment) SetAmount(value string) error {
	p.Amount = value
	return nil
}

func (p *BasePayment) SetCurrency(value string) error {
	p.Currency = value
	return nil
}

func (p *BasePayment) SetSenderReference(value string) error {
	p.Reference = value
	return nil
}

func (p *BasePayment) SetRecipientName(value string) error {
	p.Recipient = value
	return nil
}

func (p *BasePayment) SetMessage(value string) error {
	p.Msg = value
	return nil
}

func (p *BasePayment) GetErrors() map[string]error {
	return p.Errors
}
