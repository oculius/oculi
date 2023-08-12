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
		Generate() (*bytes.Buffer, errext.Error)
		FileGenerate(fileName string) errext.Error
		IOGenerate(writer io.Writer) errext.Error
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
