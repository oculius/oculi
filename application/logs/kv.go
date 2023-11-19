package logs

// KV: Key-Value Metadata
func KV(key string, value any) Metadata {
	return singleMetadata{key, value}
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
	return M(data)
}
