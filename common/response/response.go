package response

import (
	"github.com/oculius/oculi/v2/common/error-extension"
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

	Convertible interface {
		ResponseCode() int
		ResponseStatus() string
		Detail() string
		Metadata() any
	}

	DetailConvertible interface {
		Convertible
		Content() any
	}
)

func New(resp Convertible) Body {
	result := Body{
		Code:   resp.ResponseCode(),
		Status: resp.ResponseStatus(),
		Detail: resp.Detail(),
		Content: Content{
			Metadata: resp.Metadata(),
		},
	}
	normal, ok := resp.(DetailConvertible)
	err, ok2 := resp.(errext.HttpError)
	if ok && ok2 {
		panic("ambigous http response")
	} else if ok {
		result.Content.Data = normal.Content()
	} else if ok2 {
		result.Content.Error = err.Error()
	}
	return result
}
