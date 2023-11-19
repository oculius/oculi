package tk

type Kind uint

var kindString = [...]string{
	"invalid_kind", "bool",
	"int", "int8", "int16", "int32", "int64",
	"uint", "uint8", "uint16", "uint32", "uint64",
	"float32", "float64", "string", "time", "file_header",
	"uuid_string", "base36_string", "file_content_bytes", "file_content_base64",
}

const (
	InvalidKind Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	String
	Time
	FileHeader
	UUIDString
	Base36String
	FileContentBytes
	FileContentBase64
)

func (k Kind) String() string {
	return kindString[k]
}

func (k Kind) IsFromFormFile() bool {
	switch k {
	case FileHeader:
	case FileContentBytes:
	case FileContentBase64:
		return true
	}
	return false
}
