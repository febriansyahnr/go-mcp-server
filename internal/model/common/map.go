package commonModel

import "encoding/json"

type TMapAny map[string]any

func (m *TMapAny) Json() []byte {
	if m == nil {
		return []byte("{}")
	}
	if jsonByte, err := json.Marshal(m); err == nil {
		return jsonByte
	}
	return []byte("{}")
}

type TMapString map[string]string

func (m *TMapString) Json() []byte {
	if m == nil {
		return []byte("{}")
	}
	if jsonByte, err := json.Marshal(m); err == nil {
		return jsonByte
	}
	return []byte("{}")
}

func (m *TMapString) ToMapAny() TMapAny {
	if m == nil {
		return nil
	}

	result := make(TMapAny)
	for k, v := range *m {
		result[k] = v
	}
	return result
}
