package generation

import (
	"fmt"
	"net/http"

	errext "github.com/oculius/oculi/v2/common/http-error"
)

var (
	ErrFailedToGenerate = errext.NewConditional("generation:failed",
		func(i ...interface{}) string {
			return fmt.Sprintf("generation %s failed: %s", i...)
		}, func(err error) int {
			return http.StatusInternalServerError
		})
)
