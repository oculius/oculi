package json

import (
	"github.com/json-iterator/go"
	"sync"
)

type (
	Parser interface {
		Marshal(val interface{}) ([]byte, error)
		Unmarshal(data []byte, val interface{}) error
		API() jsoniter.API
	}
)

var (
	instance Parser
	once     sync.Once
)
