package response

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/paper-indonesia/pg-mcp-server/constant"
	"github.com/paper-indonesia/pg-mcp-server/pkg/util"
)

func SendSnapResponse(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(data)
}

func SendSnapResponseWithHttpCode(w http.ResponseWriter, data any, code int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	return json.NewEncoder(w).Encode(data)
}

func SnapUnauthorizedResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	var snapResponse SnapResponse

	snapResponse.ResponseCode, snapResponse.ResponseMessage = util.GenerateResponseCode(
		constant.SNAP_UNAUTHORIZED,
		"XX",
	)

	return json.NewEncoder(w).Encode(snapResponse)
}

func GetSnapResponseHttpCode(responseCode string) int {
	// convert from "2002500" to 200
	// get 3 first digit
	code := responseCode[:3]

	// convert to int
	codeInt, _ := strconv.Atoi(code)
	return codeInt
}

func SendSnapValidationResponse(w http.ResponseWriter, serviceCode string, err error) error {
	w.Header().Set("Content-Type", "application/json")

	var snapResponse SnapResponse

	var validatorErr validator.ValidationErrors
	if errors.As(err, &validatorErr) {
		// errs := make(map[string]interface{})
		var msg string
		for _, e := range validatorErr {
			if e.Tag() == "required" {
				snapResponse.ResponseCode, msg = util.GenerateResponseCode(
					constant.SNAP_INVALID_MANDATORY,
					serviceCode,
				)
				snapResponse.ResponseMessage = msg + " " + getSnapFieldFormat(e.Field())

				break
			}
			if e.Tag() == "max" || e.Tag() == "eq" || e.Tag() == "va_number" {
				snapResponse.ResponseCode, msg = util.GenerateResponseCode(
					constant.SNAP_INVALID_FIELD,
					serviceCode,
				)

				snapResponse.ResponseMessage = msg + " " + getSnapFieldFormat(e.Field())
				break
			}
		}
	}

	code := GetSnapResponseHttpCode(snapResponse.ResponseCode)
	w.WriteHeader(code)

	return json.NewEncoder(w).Encode(snapResponse)
}

func getSnapFieldFormat(field string) string {
	// set all to slice
	fieldSlice := make([]string, 0)
	for _, v := range field {
		fieldSlice = append(fieldSlice, string(v))
	}

	// get first element and set to lowercase
	fieldSlice[0] = strings.ToLower(fieldSlice[0])

	return strings.Join(fieldSlice, "")
}
