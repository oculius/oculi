package json

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/oculius/oculi/v2/common/encoding"
	"sync"
)

type JSON encoding.Encoder[jsoniter.API]

var (
	instance JSON
	once     sync.Once
)
