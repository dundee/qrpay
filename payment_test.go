package payment_test

import (
	"os"
	"testing"

	payment "github.com/dundee/go-qrcode-payment"
	"github.com/stretchr/testify/assert"
)

func TestGetQRCodeImage(t *testing.T) {
	p := payment.NewSpaydPayment()
	p.SetIBAN("CZ5855000000001265098001")

	b, _ := payment.GetQRCodeImage(p)

	// check start of PNG image
	assert.True(t, len(b) > 10)
	assert.Equal(t, uint8(0x89), b[0])
	assert.Equal(t, uint8(0x50), b[1])
}

func TestSaveQRCode(t *testing.T) {
	path := "qr.jpeg"
	defer os.Remove(path)

	p := payment.NewSpaydPayment()
	p.SetIBAN("CZ5855000000001265098001")

	payment.SaveQRCodeImageToFile(p, path)

	assert.FileExists(t, path)
}