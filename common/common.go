package common

type Payment interface {
	SetIBAN(string) error
	SetBIC(string) error
	SetAmount(value string) error
	SetCurrency(value string) error
	SetSenderReference(value string) error
	SetRecipientName(value string) error
	SetMessage(value string) error

	GetErrors() map[string]error
	GenerateString() (string, error)
}
