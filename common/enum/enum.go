package enum

import (
	"database/sql/driver"
	"github.com/oculius/oculi/v2/common/encoding/json"
	errext "github.com/oculius/oculi/v2/common/http-error"
	"net/http"
	"strings"
)

type (
	Enum interface {
		Single

		Value() (driver.Value, error)

		MarshalJSON() ([]byte, error)
	}

	Scannable interface {
		Scan(val interface{}) error
		UnmarshalJSON(val []byte) error
	}

	Single interface {
		Name() string
		Code() string
	}
)

var (
	enumCollection = map[string][]Single{}
	jsonEngine     = json.Instance()

	ErrEnumNotFound      = errext.New("enum not found", http.StatusInternalServerError)
	ErrEnumKeyRegistered = errext.New("duplicate enum key registered", http.StatusInternalServerError)
)

func Create[T Enum, V Scannable](enumKey string, enums []Single) error {
	var _ T
	var _ V

	if enumCollection[enumKey] != nil {
		return ErrEnumKeyRegistered(nil, map[string]string{
			"key": enumKey,
		})
	}

	if len(enums) == 0 || len(strings.TrimSpace(enumKey)) == 0 {
		return nil
	}
	enumCollection[enumKey] = enums
	return nil
}

func findIdx(x string, key string, selector func(e Single) string) int {
	if enumCollection[key] == nil {
		return 0
	}
	for i, v := range enumCollection[key] {
		if selector(v) == x {
			return i + 1
		}
	}
	return 0
}
