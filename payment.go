/*
Package for creating QR codes for payments.

[Short Payment Descriptor](https://en.wikipedia.org/wiki/Short_Payment_Descriptor) format and
[EPC QR Code](https://en.wikipedia.org/wiki/EPC_QR_code) (SEPA) format is supported.

## Usage

### Generating QR code image for Short Payment Descriptor format

	import payment "github.com/dundee/go-qrcode-payment"

	p := payment.NewSpaydPayment()
	p.SetIBAN("CZ5855000000001265098001")
	p.SetAmount("10.8")
	p.SetDate(time.Date(2021, 12, 24, 0, 0, 0, 0, time.UTC))
	p.SetMessage("M")
	p.SetRecipientName("go")
	p.SetNofificationType('E')
	p.SetNotificationValue("daniel@milde.cz")
	p.SetExtendedAttribute("vs", "1234567890")

	payment.SaveQRCodeImageToFile(p, "qr-payment.png")

### Generating QR code image for EPC QR Code

	import payment "github.com/dundee/go-qrcode-payment"

	p := payment.NewEpcPayment()
	p.SetIBAN("CZ5855000000001265098001")
	p.SetAmount("10.8")
	p.SetMessage("M")
	p.SetRecipientName("go")

	payment.SaveQRCodeImageToFile(p, "qr-payment.png")

QR code image encoding uses [skip2/go-qrcode](https://github.com/skip2/go-qrcode).

### Getting QR code content for Short Payment Descriptor format

	import payment "github.com/dundee/go-qrcode-payment"

	p := payment.NewSpaydPayment()
	p.SetIBAN("CZ5855000000001265098001")
	p.SetAmount("108")

	fmt.Println(payment.GenerateString())
	// Output: SPD*1.0*ACC:CZ5855000000001265098001*AM:108*

*/
package payment

import (
	"github.com/dundee/go-qrcode-payment/common"
	"github.com/dundee/go-qrcode-payment/epc"
	"github.com/dundee/go-qrcode-payment/spayd"

	qrcode "github.com/skip2/go-qrcode"
)

func NewSpaydPayment() *spayd.SpaydPayment {
	return spayd.NewSpaydPayment()
}

func NewEpcPayment() *epc.EpcPayment {
	return epc.NewEpcPayment()
}

func SaveQRCodeImageToFile(payment common.QRCodeGenerator, path string) error {
	content, err := payment.GenerateString()
	if err != nil {
		return err
	}
	return qrcode.WriteFile(content, qrcode.Medium, 400, path)
}

func GetQRCodeImage(payment common.QRCodeGenerator) ([]byte, error) {
	content, err := payment.GenerateString()
	if err != nil {
		return nil, err
	}
	return qrcode.Encode(content, qrcode.Medium, 400)
}
