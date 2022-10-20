package logs

type (
	Info interface {
		Message() string
		Metadata() map[string]any
		EditMetadata(key string, value any)
		RemoveMetadata(key string)
		ClearMetadata()
	}

	Metadata interface {
		Apply(si Info)
	}

	info struct {
		metadata map[string]interface{}
		msg      string
	}

	singleMetadata struct {
		key   string
		value any
	}

	groupMetadata map[string]any
)

// KV: Key-Value Metadata
func KV(key string, value any) Metadata {
	return singleMetadata{key, value}
}

func (sm singleMetadata) Apply(i Info) {
	i.EditMetadata(sm.key, sm.value)
}

// M: Map Metadata
func M(metadata map[string]any) Metadata {
	return groupMetadata(metadata)
}

func (gm groupMetadata) Apply(i Info) {
	for key, value := range gm {
		i.EditMetadata(key, value)
	}
}

func (i *info) Message() string {
	return i.msg
}

func (i *info) Metadata() map[string]any {
	return i.metadata
}

func (i *info) EditMetadata(key string, val any) {
	i.metadata[key] = val
}

func (i *info) RemoveMetadata(key string) {
	delete(i.metadata, key)
}

func (i *info) ClearMetadata() {
	for key := range i.metadata {
		delete(i.metadata, key)
	}
}

func NewInfo(msg string, metadata ...Metadata) Info {
	var result Info = &info{
		msg:      msg,
		metadata: map[string]any{},
	}
	for _, md := range metadata {
		md.Apply(result)
	}
	return result
}
