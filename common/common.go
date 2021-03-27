package common

type QRCodeGenerator interface {
	GenerateString() (string, error)
}
