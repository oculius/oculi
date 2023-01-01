package json

import "sync"

type (
	Engine interface {
		Marshal(val interface{}) ([]byte, error)
		Unmarshal(data []byte, val interface{}) error
	}
)

var (
	instance Engine
	once     sync.Once
)
