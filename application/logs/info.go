package logs

type (
	info struct {
		metadata map[string]interface{}
		msg      string
	}
)

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

func (i *info) Message() string {
	return i.msg
}

func (i *info) Metadata() map[string]any {
	return i.metadata
}

func (i *info) EditMetadata(key string, val any) {
	if len(key) == 0 {
		return
	}

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
