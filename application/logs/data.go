package logs

type (
	singleMetadata struct {
		key   string
		value any
	}

	M map[string]any
)

func (sm singleMetadata) Apply(i Info) {
	i.EditMetadata(sm.key, sm.value)
}

func (m M) Apply(i Info) {
	for key, value := range m {
		i.EditMetadata(key, value)
	}
}
