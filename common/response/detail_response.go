package response

import "net/http"

type (
	detailResponse struct {
		detail   string
		data     any
		metadata any
	}
)

func NewResponse(detail string, data any, metadata any) DetailConvertible {
	return &detailResponse{detail, data, metadata}
}

func (o *detailResponse) ResponseCode() int {
	return http.StatusOK
}

func (o *detailResponse) ResponseStatus() string {
	return http.StatusText(http.StatusOK)
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
