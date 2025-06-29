package response

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type SnapResponse struct {
	ResponseCode       string         `json:"responseCode" example:"200xx200"`
	ResponseMessage    string         `json:"responseMessage" example:"Success"`
	VirtualAccountData any            `json:"virtualAccountData,omitempty"`
	AdditionalInfo     map[string]any `json:"additionalInfo,omitempty"`
}

func (r *SnapResponse) String() string {
	jsonStr, err := json.Marshal(r)
	if err != nil {
		return fmt.Sprintf(`{"error": "%s"}`, err.Error())
	}
	return string(jsonStr)
}

type ErrorDetail struct {
	Type    string      `json:"type,omitempty"`
	Message interface{} `json:"message,omitempty"`
	TraceId string      `json:"trace_id,omitempty"`
}
