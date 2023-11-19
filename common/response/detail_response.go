package response

import "net/http"

type (
	detailResponse struct {
		detail   string
		data     any
		metadata any
		status   int
	}
)

func NewResponse(detail string, data any, metadata any) DetailConvertible {
	return &detailResponse{detail, data, metadata, http.StatusOK}
}

func NewResponseWithStatus(detail string, data any, metadata any, status int) DetailConvertible {
	return &detailResponse{detail, data, metadata, status}
}

func (o *detailResponse) ResponseCode() int {
	return o.status
}

func (o *detailResponse) ResponseStatus() string {
	return http.StatusText(o.status)
}

func (o *detailResponse) Detail() string {
	return o.detail
}

func (o *detailResponse) Metadata() any {
	return o.metadata
}

func (o *detailResponse) Content() any {
	return o.data
}
