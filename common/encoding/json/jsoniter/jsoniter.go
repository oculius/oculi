package jsoniter

import (
	jlib "github.com/json-iterator/go"
	"github.com/oculius/oculi/v2/common/encoding/json"
)

type (
	jsoniter struct {
		json jlib.API
	}
)

var instance json.Engine = &jsoniter{
	json: jlib.ConfigCompatibleWithStandardLibrary,
}

func (i *jsoniter) Marshal(val interface{}) ([]byte, error) {
	return i.json.Marshal(val)
}

func (i *jsoniter) Unmarshal(data []byte, val interface{}) error {
	return i.json.Unmarshal(data, val)
}

func New() json.Engine {
	return instance
}
