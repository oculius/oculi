package generation

import (
	"bytes"
	"io"

	errext "github.com/oculius/oculi/v2/common/http-error"
)

type (
	Generator interface {
		Generate() (*bytes.Buffer, errext.HttpError)
		FileGenerate(fileName string) errext.HttpError
		IOGenerate(writer io.Writer) errext.HttpError
	}
)
