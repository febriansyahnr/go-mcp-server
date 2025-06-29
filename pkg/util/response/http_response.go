package response

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	pdkConst "github.com/paper-indonesia/pdk/v2/constant"
	"github.com/paper-indonesia/pg-mcp-server/pkg/dictionary"
	pkgErrors "github.com/paper-indonesia/pg-mcp-server/pkg/error"
	"github.com/paper-indonesia/pg-mcp-server/pkg/util"
	"github.com/paper-indonesia/pg-mcp-server/pkg/validatorExt"
)

// SendResponseOK send json response to client. If withPredefineResponse=true, it will send with
// predefine struct (code, message, data).
func SendResponseOK(w http.ResponseWriter, data interface{}, withPredefineResponse bool) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if withPredefineResponse {
		resp := Response{
			Code:    GetHttpCodeService(HttpStatusOK),
			Message: http.StatusText(http.StatusOK),
			Data:    data,
		}

		return json.NewEncoder(w).Encode(resp)
	}
	return json.NewEncoder(w).Encode(data)
}

func SendResponseError(ctx context.Context, w http.ResponseWriter, errMessage error, data interface{}) error {
	errCode, errObj := pkgErrors.ExtractError(errMessage)
	errType := pkgErrors.GetErrorType(errCode)
	code, statusCode := GetSnapHttpAndResponseCode(data, errCode)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorDetail := ErrorDetail{
		Type:    errType,
		Message: errObj.Error(),
	}

	if traceID, ok := ctx.Value(pdkConst.CtxTraceIdKey).(string); ok {
		errorDetail.TraceId = traceID
	}

	var validatorError validator.ValidationErrors
	if errors.As(errMessage, &validatorError) {
		errs := make(map[string]interface{})
		for _, e := range validatorError {
			errs[e.Field()] = e.Translate(validatorExt.GetTranslator())
		}

		errorDetail.Message = errs
	}

	resp := Response{
		Code:    GetHttpCodeServiceError(code, errCode),
		Message: dictionary.Dict.SetDictionaryMessage(ctx, errCode),
		Error:   errorDetail,
		Data:    data,
	}

	if errCode == pkgErrors.ErrCodeResourceNotFound {
		resp.Message = errObj.Error()
	}

	return json.NewEncoder(w).Encode(resp)
}

func SendResponseCreated(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	resp := Response{
		Code:    GetHttpCodeService(HttpStatusCreated),
		Data:    data,
		Message: "Created",
	}

	err := json.NewEncoder(w).Encode(resp)
	return err
}

func GetSnapHttpAndResponseCode(data any, errCode string) (code string, httpCode int) {
	var snapResponse SnapResponse

	dataB, err := json.Marshal(data)
	err = json.Unmarshal(dataB, &snapResponse)
	if err == nil {
		code, httpCode = extractStatusFromResponseCode(errCode, snapResponse.ResponseCode)
	} else {
		code, httpCode = GetHTTPStatusCode(errCode)
	}

	if util.InArray(httpCode, []int{
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
	}) {
		httpCode = http.StatusRequestTimeout
	}
	return code, httpCode
}

func SendResponse(w http.ResponseWriter, data interface{}, code int, withPredefineResponse bool) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if withPredefineResponse {
		resp := Response{
			Code:    GetHttpCodeService(HttpStatusOK),
			Message: http.StatusText(http.StatusOK),
			Data:    data,
		}

		return json.NewEncoder(w).Encode(resp)
	}
	return json.NewEncoder(w).Encode(data)
}
