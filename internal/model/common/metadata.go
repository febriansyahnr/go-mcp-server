package commonModel

import "encoding/json"

type Metadata map[string]any

func NewMetadata() Metadata {
	return make(Metadata)
}

func (m *Metadata) Set(key string, value any) {
	(*m)[key] = value
}

func (m *Metadata) Get(key string) any {
	if val, ok := (*m)[key]; ok {
		return val
	}
	return nil
}

func (m *Metadata) Del(key string) {
	delete(*m, key)
}

func (m *Metadata) Json() []byte {
	json, _ := json.Marshal(m)
	return json
}

func (m *Metadata) FromStr(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), m)
}
