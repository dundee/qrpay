package base

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
	Errors    map[string]error
}

func NewPayment() *Payment {
	return &Payment{
		Errors: make(map[string]error),
	}
}

func (p *Payment) SetIBAN(value string) error {
	_, err := iban.Parse(value)
	if err != nil {
		p.Errors["iban"] = err
	} else {
		p.IBAN = value
	}
	return err
}

func (p *Payment) SetBIC(value string) error {
	_, err := swift.Parse(value)
	if err != nil {
		p.Errors["bic"] = err
	} else {
		p.BIC = value
	}
	return err
}

func (p *Payment) SetAmount(value string) error {
	p.Amount = value
	return nil
}

func (p *Payment) SetCurrency(value string) error {
	p.Currency = value
	return nil
}

func (p *Payment) SetSenderReference(value string) error {
	p.Reference = value
	return nil
}

func (p *Payment) SetRecipientName(value string) error {
	p.Recipient = value
	return nil
}

func (p *Payment) SetMessage(value string) error {
	p.Msg = value
	return nil
}

func (p *Payment) GetErrors() map[string]error {
	return p.Errors
}
