package json

type (
	Marshaller interface {
		Marshal(val interface{}) ([]byte, error)
	}

	Unmarshaller interface {
		Unmarshal(data []byte, val interface{}) error
	}

	Engine interface {
		Marshaller
		Unmarshaller
	}
)
