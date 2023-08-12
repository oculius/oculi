package enum

type (
	defaultSingle struct {
		name string
		code string
	}
)

func New(name, code string) Single {
	return defaultSingle{name, code}
}

func (d defaultSingle) Name() string {
	return d.name
}

func (d defaultSingle) Code() string {
	return d.code
}
