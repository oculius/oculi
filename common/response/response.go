package response

import (
	"github.com/oculius/oculi/v2/common/error"
)

type (
	Response struct {
		Code   int          `json:"code"`
		Status string       `json:"status"`
		Detail string       `json:"detail"`
		Data   ResponseData `json:"data,omitempty"`
	}

	ResponseData struct {
		Error    string `json:"error,omitempty"`
		Content  any    `json:"content,omitempty"`
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

func New(resp HttpResponse) Response {
	result := Response{
		Code:   resp.ResponseCode(),
		Status: resp.ResponseStatus(),
		Detail: resp.Detail(),
		Data: ResponseData{
			Metadata: resp.Metadata(),
		},
	}
	normal, ok := resp.(OkHttpResponse)
	err, ok2 := resp.(gerr.Error)
	if ok && ok2 {
		panic("ambigous http response")
	} else if ok {
		result.Data.Content = normal.Content()
	} else if ok2 {
		result.Data.Error = err.Error()
	}
	return result
}
