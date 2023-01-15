package json

import (
	jsoniter "github.com/json-iterator/go"
)

type (
	impl struct {
		json jsoniter.API
	}
)

func (i *impl) API() jsoniter.API {
	return i.json
}

func (i *impl) Marshal(val interface{}) ([]byte, error) {
	return i.json.Marshal(val)
}

func (i *impl) Unmarshal(data []byte, val interface{}) error {
	return i.json.Unmarshal(data, val)
}

func New() Parser {
	once.Do(func() {
		if instance != nil {
			panic("jsoniter: singleton instance error")
		}
		instance = &impl{
			json: jsoniter.ConfigCompatibleWithStandardLibrary,
		}
	})
	return instance
}
