package payment

import (
	"github.com/dundee/go-qrcode-payment/common"
	"github.com/dundee/go-qrcode-payment/spayd"

	qrcode "github.com/skip2/go-qrcode"
)

func NewSpaydPayment() *spayd.SpaydPayment {
	return spayd.NewSpaydPayment()
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
