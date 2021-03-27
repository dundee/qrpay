package main

import (
	"fmt"
	"time"

	payment "github.com/dundee/go-qrcode-payment"
)

func main() {
	p := payment.NewSpaydPayment()

	p.SetIBAN("CZ5855000000001265098001")
	p.SetAmount("100")
	p.SetDate(time.Date(2021, 12, 24, 0, 0, 0, 0, time.UTC))
	p.SetMessage("M")
	p.SetRecipientName("go")
	p.SetNofificationType('E')
	p.SetNotificationValue("daniel@milde.cz")
	p.SetExtendedAttribute("vs", "1234567890")

	if err := payment.SaveQRCodeImageToFile(p, "qr-payment.png"); err != nil {
		fmt.Printf("could not generate payment QR code: %v", err)
	}

}
