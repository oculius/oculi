package enum

import "database/sql/driver"

type BaseEnum int

func (b BaseEnum) AbstractMarshalJSON(enum Single) ([]byte, error) {
	return jsonEngine.Marshal(enum.Name())
}

func (b BaseEnum) AbstractUnmarshalJSON(val []byte, key string) (int, error) {
	var rawValue string
	if err := jsonEngine.Unmarshal(val, &rawValue); err != nil {
		return 0, err
	}

	idx := findIdx(rawValue, key, func(e Single) string { return e.Name() })
	if idx == 0 {
		return 0, ErrEnumNotFound(nil, nil)
	}
	return idx, nil
}

func (b BaseEnum) AbstractScan(val interface{}, key string) (int, error) {
	var dbValue string
	switch v := val.(type) {
	case string:
		dbValue = v
	case []byte:
		dbValue = string(v)
	default:
		return 0, ErrEnumNotFound(nil, nil)
	}
	idx := findIdx(dbValue, key, func(e Single) string { return e.Code() })
	if idx == 0 {
		return 0, ErrEnumNotFound(nil, nil)
	}
	return idx, nil
}

func (b BaseEnum) AbstractValue(enum Single) (driver.Value, error) {
	return enum.Code(), nil
}

func (b BaseEnum) AbstractName(key string) string {
	arr, ok := enumCollection[key]
	if !ok {
		return ""
	}
	if int(b) < 1 || int(b) > len(arr) {
		return ""
	}
	return arr[int(b)-1].Name()
}

func (b BaseEnum) AbstractCode(key string) string {
	arr, ok := enumCollection[key]
	if !ok {
		return ""
	}
	if int(b) < 1 || int(b) > len(arr) {
		return ""
	}
	return arr[int(b)-1].Code()
}
