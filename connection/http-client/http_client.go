package httpclient

import (
	"fmt"
	errext "github.com/oculius/oculi/v2/common/http-error"
	"io"
	"net/http"
	"time"
)

type (
	WebClientFactory interface {
		Create(timeout time.Duration) WebClient
		CreateWithRetry(timeout time.Duration, retryCount int) WebClient
	}

	WebClient interface {
		Get(url string, headers http.Header, queryString map[string]string) (*http.Response, error)
		Post(url string, body io.Reader, headers http.Header) (*http.Response, error)
		Put(url string, body io.Reader, headers http.Header) (*http.Response, error)
		Patch(url string, body io.Reader, headers http.Header) (*http.Response, error)
		Delete(url string, headers http.Header) (*http.Response, error)
		Do(r *http.Request) (*http.Response, error)
	}
)

var (
	ErrWebClient = errext.NewConditional("webclient:generic_error",
		func(i ...interface{}) string {
			return fmt.Sprintf("web client error: %s", i...)
		},
		func(error) int {
			return http.StatusBadGateway
		})
)
