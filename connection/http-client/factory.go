package httpclient

import (
	"time"

	"github.com/gojek/heimdall/httpclient"
)

type (
	clientFactory struct{}
)

func (cf *clientFactory) Create(timeout time.Duration) WebClient {
	return &client{
		Heimdall: *httpclient.NewClient(httpclient.WithHTTPTimeout(timeout)),
	}
}

func (cf *clientFactory) CreateWithRetry(timeout time.Duration, retryCount int) WebClient {
	return &client{
		Heimdall: *httpclient.NewClient(
			httpclient.WithHTTPTimeout(timeout),
			httpclient.WithRetryCount(retryCount),
		),
	}
}

func New() WebClientFactory {
	return &clientFactory{}
}
