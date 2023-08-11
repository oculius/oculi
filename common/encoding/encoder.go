package encoding

type (
	Encoder[T any] interface {
		Marshal(val interface{}) ([]byte, error)
		Unmarshal(data []byte, val interface{}) error
		API() T
	}

	GeneralEncoder Encoder[any]
)
