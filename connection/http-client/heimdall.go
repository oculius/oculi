package httpclient

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	neturl "net/url"

	heimdall "github.com/gojek/heimdall/httpclient"
)

type (
	client struct {
		Heimdall heimdall.Client
	}

	WebClientErrorDetails struct {
		ResponseStatus     string `json:"response_status"`
		ResponseStatusCode int    `json:"response_status_code"`
	}
)

func (c *client) Get(url string, headers http.Header, queryString map[string]string) (*http.Response, error) {

	if len(queryString) > 0 {
		format := "?%s=%s"
		for key, value := range queryString {
			url += fmt.Sprintf(format, key, neturl.QueryEscape(value))
			format = "&%s=%s"
		}
	}

	response, err := c.Heimdall.Get(url, headers)
	return c.checkClientError(response, err)
}

func (c *client) Post(url string, body io.Reader, headers http.Header) (*http.Response, error) {

	response, err := c.Heimdall.Post(url, body, headers)
	return c.checkClientError(response, err)
}

func (c *client) Put(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	response, err := c.Heimdall.Put(url, body, headers)
	return c.checkClientError(response, err)
}

func (c *client) Patch(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	response, err := c.Heimdall.Patch(url, body, headers)
	return c.checkClientError(response, err)
}

func (c *client) Delete(url string, headers http.Header) (*http.Response, error) {
	response, err := c.Heimdall.Delete(url, headers)
	return c.checkClientError(response, err)
}

func (c *client) Do(r *http.Request) (*http.Response, error) {
	response, err := c.Heimdall.Do(r)
	return c.checkClientError(response, err)
}

func (c *client) checkClientError(response *http.Response, err error) (*http.Response, error) {
	if err != nil {
		return response, ErrWebClient(err, nil, err.Error())
	} else if response.StatusCode >= http.StatusBadRequest {
		err = errors.New("bad status")
		return response, ErrWebClient(err,
			WebClientErrorDetails{
				ResponseStatus:     response.Status,
				ResponseStatusCode: response.StatusCode,
			}, err)
	}
	return response, nil
}
