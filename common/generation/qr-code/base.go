package qrcode

type (
	qrcodeGenerator struct {
		Content              string
		ErrorCorrectionLevel ErrorCorrectionLevel
		Encoding             Encoding
		Size                 int
	}
)
