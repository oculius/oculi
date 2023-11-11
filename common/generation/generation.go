package generation

import (
	"bytes"
	"fmt"
	errext "github.com/oculius/oculi/v2/common/error-extension"
	"io"
	"net/http"
)

type (
	Generator interface {
		Generate() (*bytes.Buffer, errext.HttpError)
		FileGenerate(fileName string) errext.HttpError
		IOGenerate(writer io.Writer) errext.HttpError
	}
)

var (
	ErrFailedToGenerate = errext.NewConditional("generation:failed",
		func(i ...interface{}) string {
			return fmt.Sprintf("generation %s failed: %s", i...)
		}, func(err error) int {
			return http.StatusInternalServerError
		})
)
