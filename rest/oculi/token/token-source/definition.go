package ts

type (
	TokenSource uint
)

var tokenSourceString = []string{"invalid_source", "query", "parameter", "header", "cookie", "form", "form_file"}

const (
	InvalidSource TokenSource = iota
	Query
	Parameter
	Header
	Cookie
	Form
	FormFile
)

func (ts TokenSource) String() string {
	return tokenSourceString[ts]
}
