package response

import (
	"github.com/oculius/oculi/v2/common/error"
)

type (
	Body struct {
		Code    int     `json:"code"`
		Status  string  `json:"status"`
		Detail  string  `json:"detail"`
		Content Content `json:"content,omitempty"`
	}

	Content struct {
		Error    string `json:"error,omitempty"`
		Data     any    `json:"data,omitempty"`
		Metadata any    `json:"metadata,omitempty"`
	}

	HttpResponse interface {
		ResponseCode() int
		ResponseStatus() string
		Detail() string
		Metadata() any
	}

	OkHttpResponse interface {
		HttpResponse
		Content() any
	}
)

func New(resp HttpResponse) Body {
	result := Body{
		Code:   resp.ResponseCode(),
		Status: resp.ResponseStatus(),
		Detail: resp.Detail(),
		Content: Content{
			Metadata: resp.Metadata(),
		},
	}
	normal, ok := resp.(OkHttpResponse)
	err, ok2 := resp.(gerr.Error)
	if ok && ok2 {
		panic("ambigous http response")
	} else if ok {
		result.Content.Data = normal.Content()
	} else if ok2 {
		result.Content.Error = err.Error()
	}
	return result
}
