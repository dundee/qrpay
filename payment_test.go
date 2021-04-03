package qrpay_test

import (
	"os"
	"testing"

	"github.com/dundee/qrpay"
	"github.com/stretchr/testify/assert"
)

func TestGetQRCodeImage(t *testing.T) {
	p := qrpay.NewSpaydPayment()
	p.SetIBAN("CZ5855000000001265098001")

	b, _ := qrpay.GetQRCodeImage(p)

	// check start of PNG image
	assert.True(t, len(b) > 10)
	assert.Equal(t, uint8(0x89), b[0])
	assert.Equal(t, uint8(0x50), b[1])
}

func TestGetQRCodeImageWithErr(t *testing.T) {
	p := qrpay.NewSpaydPayment()

	b, err := qrpay.GetQRCodeImage(p)

	assert.Nil(t, b)
	assert.Equal(t, "IBAN is mandatory", err.Error())
}

func TestSaveQRCode(t *testing.T) {
	path := "qr.jpeg"
	defer os.Remove(path)

	p := qrpay.NewSpaydPayment()
	p.SetIBAN("CZ5855000000001265098001")

	qrpay.SaveQRCodeImageToFile(p, path)

	assert.FileExists(t, path)
}

func TestSaveEpcQRCode(t *testing.T) {
	path := "qr.jpeg"
	defer os.Remove(path)

	p := qrpay.NewEpcPayment()
	p.SetRecipientName("Red Cross")
	p.SetIBAN("CZ5855000000001265098001")

	qrpay.SaveQRCodeImageToFile(p, path)

	assert.FileExists(t, path)
}

func TestSaveQRCodeWithErr(t *testing.T) {
	path := "qr.jpeg"
	p := qrpay.NewSpaydPayment()
	err := qrpay.SaveQRCodeImageToFile(p, path)

	assert.Equal(t, "IBAN is mandatory", err.Error())
}
