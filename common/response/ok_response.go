package response

import "net/http"

type (
	okResponse struct {
		detail   string
		data     any
		metadata any
	}
)

func NewOkResponse(detail string, data any, metadata any) OkHttpResponse {
	return &okResponse{detail, data, metadata}
}

func (o *okResponse) ResponseCode() int {
	return http.StatusOK
}

func (o *okResponse) ResponseStatus() string {
	return http.StatusText(http.StatusOK)
}

func (o *okResponse) Detail() string {
	return o.detail
}

func (o *okResponse) Metadata() any {
	return o.metadata
}

func (o *okResponse) Content() any {
	return o.data
}
