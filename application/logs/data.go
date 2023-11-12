package logs

import "fmt"

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

func anyToString(item any) string {
	key, ok := item.(string)
	if ok {
		return key
	}

	stringer, ok := item.(fmt.Stringer)
	if ok {
		return stringer.String()
	}

	return fmt.Sprintf("%#v", item)
}

// KVs: Multiple Key-Value Metadata
func KVs(args ...any) Metadata {
	if len(args) == 0 {
		return singleMetadata{}
	}

	if len(args) < 2 {
		return singleMetadata{anyToString(args[0]), ""}
	}

	data := make(map[string]any, len(args)/2)
	for i := 0; i < len(args); i += 2 {
		key := anyToString(args[i])
		data[key] = args[i+1]
	}
	return groupMetadata(data)
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
