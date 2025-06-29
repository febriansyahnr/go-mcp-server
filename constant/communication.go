package constant

import (
	"encoding/json"
	"strconv"
)

type EmailPriority int

const (
	EmailPrioritL0 EmailPriority = iota
	EmailPrioritL1
	EmailPrioritL2
)

func (e EmailPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.Itoa(int(e)))
}
