/*
Package for creating QR codes for payments.

Short Payment Descriptor format and EPC QR Code (SEPA) format is supported.

- Generating QR code image for Short Payment Descriptor format

	import "github.com/dundee/qrpay"

	p := qrpay.NewSpaydPayment()
	p.SetIBAN("CZ5855000000001265098001")
	p.SetAmount("10.8")
	p.SetDate(time.Date(2021, 12, 24, 0, 0, 0, 0, time.UTC))
	p.SetMessage("M")
	p.SetRecipientName("go")
	p.SetNofificationType('E')
	p.SetNotificationValue("daniel@milde.cz")
	p.SetExtendedAttribute("vs", "1234567890")

	qrpay.SaveQRCodeImageToFile(p, "qr-payment.png")

- Generating QR code image for EPC QR Code

	import "github.com/dundee/qrpay"

	p := qrpay.NewEpcPayment()
	p.SetIBAN("CZ5855000000001265098001")
	p.SetAmount("10.8")
	p.SetMessage("M")
	p.SetRecipientName("go")

	qrpay.SaveQRCodeImageToFile(p, "qr-payment.png")

QR code image encoding uses https://github.com/skip2/go-qrcode

- Getting QR code content for Short Payment Descriptor format

	import "github.com/dundee/qrpay"

	p := qrpay.NewSpaydPayment()
	p.SetIBAN("CZ5855000000001265098001")
	p.SetAmount("108")

	fmt.Println(qrpay.GenerateString())
	// Output: SPD*1.0*ACC:CZ5855000000001265098001*AM:108*

*/
package qrpay

import (
	"github.com/dundee/qrpay/epc"
	"github.com/dundee/qrpay/spayd"

	qrcode "github.com/skip2/go-qrcode"
)

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

func NewSpaydPayment() *spayd.SpaydPayment {
	return spayd.NewSpaydPayment()
}

func NewEpcPayment() *epc.EpcPayment {
	return epc.NewEpcPayment()
}

func SaveQRCodeImageToFile(payment Payment, path string) error {
	content, err := payment.GenerateString()
	if err != nil {
		return err
	}
	return qrcode.WriteFile(content, qrcode.Medium, 400, path)
}

func GetQRCodeImage(payment Payment) ([]byte, error) {
	content, err := payment.GenerateString()
	if err != nil {
		return nil, err
	}
	return qrcode.Encode(content, qrcode.Medium, 400)
}

// check validity
var _ Payment = (*spayd.SpaydPayment)(nil)
var _ Payment = (*epc.EpcPayment)(nil)
