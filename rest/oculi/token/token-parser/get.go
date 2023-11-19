package tp

import (
	tk "github.com/oculius/oculi/v2/rest/oculi/token/token-kind"
)

func Get(kind tk.Kind) (Parser, bool) {
	var parser Parser
	switch kind {
	case tk.Bool:
		parser = BoolParser()
	case tk.Int:
		parser = IntParser()
	case tk.Int8:
		parser = Int8Parser()
	case tk.Int16:
		parser = Int16Parser()
	case tk.Int32:
		parser = Int32Parser()
	case tk.Int64:
		parser = Int64Parser()
	case tk.Uint:
		parser = UintParser()
	case tk.Uint8:
		parser = Uint8Parser()
	case tk.Uint16:
		parser = Uint16Parser()
	case tk.Uint32:
		parser = Uint32Parser()
	case tk.Uint64:
		parser = Uint64Parser()
	case tk.Float32:
		parser = Float32Parser()
	case tk.Float64:
		parser = Float64Parser()
	case tk.String:
		parser = StringParser()
	case tk.Time:
		parser = TimeParser()
	case tk.FileHeader:
		parser = FileHeaderParser()
	case tk.UUIDString:
		parser = UUIDStringParser()
	case tk.Base36String:
		return nil, false
		// panic("parser is not implemented yet")
	case tk.FileContentBytes:
		parser = FileContentBytesParser()
	case tk.FileContentBase64:
		parser = FileContentBase64Parser()
	default:
		return nil, false
	}
	return parser, true
}
