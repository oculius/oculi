package tp

import (
	"encoding/base64"
	gerr "github.com/oculius/oculi/v2/common/error"
	"io"
	"mime/multipart"
)

type (
	fileHeaderParser struct{}

	fcbytesParser struct{}

	fcb64Parser struct{}
)

func fileReader(val *multipart.FileHeader, t Token) ([]byte, gerr.Error) {
	file, err := val.Open()
	if err != nil {
		return nil, ErrFormFile(err, t.Metadata())
	}
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, ErrFormFile(err, t.Metadata())
	}
	return content, nil
}

func (f fcb64Parser) Parse(t Token, value any) (any, gerr.Error) {
	val, ok := value.(*multipart.FileHeader)
	if !ok {
		return nil, ErrInvalidInputValueFileHeader
	}

	content, err := fileReader(val, t)
	if err != nil {
		return nil, err
	}

	b64 := base64.StdEncoding.EncodeToString(content)
	return b64, nil
}

func (f fcbytesParser) Parse(t Token, value any) (any, gerr.Error) {
	val, ok := value.(*multipart.FileHeader)
	if !ok {
		return nil, ErrInvalidInputValueFileHeader
	}

	content, err := fileReader(val, t)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (f fileHeaderParser) Parse(_ Token, value any) (any, gerr.Error) {
	val, ok := value.(*multipart.FileHeader)
	if !ok {
		return nil, ErrInvalidInputValueFileHeader
	}
	return val, nil
}

func FileHeaderParser() Parser {
	return fileHeaderParser{}
}

func FileContentBytesParser() Parser {
	return fcbytesParser{}
}

func FileContentBase64Parser() Parser {
	return fcb64Parser{}
}
