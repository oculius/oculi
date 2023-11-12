package qrcode

import (
	"bytes"
	"image/png"
	"io"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	generation2 "github.com/oculius/oculi/v2/application/generation"
	"github.com/oculius/oculi/v2/common/http-error"
)

type (
	qrcodeGenerator struct {
		Content              string
		ErrorCorrectionLevel ErrorCorrectionLevel
		Encoding             Encoding
		Size                 int
	}
)

func New(content string, opts ...QRCodeOption) generation2.Generator {
	result := &qrcodeGenerator{
		Content:              content,
		Size:                 200,
		Encoding:             Auto,
		ErrorCorrectionLevel: MediumHighCorrection,
	}
	for _, opt := range opts {
		opt.Apply(result)
	}
	return result
}

func (q *qrcodeGenerator) Generate() (*bytes.Buffer, httperror.HttpError) {
	result := new(bytes.Buffer)
	if err := q.IOGenerate(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (q *qrcodeGenerator) FileGenerate(fileName string) httperror.HttpError {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return generation2.ErrFailedToGenerate(err, nil, "qrcode", err.Error())
	}

	if err = q.IOGenerate(file); err != nil {
		return generation2.ErrFailedToGenerate(err, nil, "qrcode", err.Error())
	}

	if err2 := file.Close(); err2 != nil {
		return generation2.ErrFailedToGenerate(err2, nil, "qrcode", err2.Error())
	}
	return nil
}

func (q *qrcodeGenerator) IOGenerate(writer io.Writer) httperror.HttpError {
	qrImgData, err := q.process()
	if err != nil {
		return generation2.ErrFailedToGenerate(err, nil, "qrcode", err.Error())
	}

	if err = png.Encode(writer, qrImgData); err != nil {
		return generation2.ErrFailedToGenerate(err, nil, "qrcode", err.Error())
	}
	return nil
}

func (q *qrcodeGenerator) process() (barcode.Barcode, error) {
	qrCode, err := qr.Encode(q.Content, q.ErrorCorrectionLevel.real(), q.Encoding.real())
	if err != nil {
		return nil, err
	}

	qrCodeResult, err := barcode.Scale(qrCode, q.Size, q.Size)
	if err != nil {
		return nil, err
	}

	return qrCodeResult, nil
}
