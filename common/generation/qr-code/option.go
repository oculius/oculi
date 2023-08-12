package qrcode

type (
	QRCodeOption interface {
		Apply(qrcode *qrcodeGenerator)
	}

	withSize               int
	withEncoding           struct{ Encoding }
	withErrCorrectionLevel struct{ ErrorCorrectionLevel }
)

func WithSize(size int) QRCodeOption {
	return withSize(size)
}

func WithEncoding(enc Encoding) QRCodeOption {
	return withEncoding{enc}
}

func WithErrCorrectionLevel(ecl ErrorCorrectionLevel) QRCodeOption {
	return withErrCorrectionLevel{ecl}
}

func (w withSize) Apply(qrcode *qrcodeGenerator) {
	qrcode.Size = int(w)
}

func (w withEncoding) Apply(qrcode *qrcodeGenerator) {
	qrcode.Encoding = w.Encoding
}

func (w withErrCorrectionLevel) Apply(qrcode *qrcodeGenerator) {
	qrcode.ErrorCorrectionLevel = w.ErrorCorrectionLevel
}
